package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

var emptyResponse = new(string)

type Option func(api *Api)

func EnableDebug(api *Api) {
	api.debug = true
}

type Api struct {
	hc                *http.Client
	address           string
	debug             bool
	common            *service
	Auth              *Auth
	App               *App
	Log               *Log
	Sync              *Sync
	TransferInfo      *TransferInfo
	TorrentManagement *TorrentManagement
	Rss               *Rss
	Search            *Search
}

type service struct {
	api *Api
}

func NewApi(address string, options ...Option) (api *Api, err error) {
	api = &Api{}
	address = strings.TrimSuffix(address, "/")
	api.address = address
	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}
	api.hc = &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	api.common = &service{
		api: api,
	}

	for _, opt := range options {
		opt(api)
	}

	api.Auth = (*Auth)(api.common)
	api.App = (*App)(api.common)
	api.Log = (*Log)(api.common)
	api.Sync = (*Sync)(api.common)
	api.TransferInfo = (*TransferInfo)(api.common)
	api.TorrentManagement = (*TorrentManagement)(api.common)
	api.Rss = (*Rss)(api.common)
	api.Search = (*Search)(api.common)

	return api, nil
}

func (a *Api) doRequest(ctx context.Context, method, path string, queryParams, formData url.Values, v any) (err error) {
	var link = fmt.Sprintf("%s%s", a.address, path)
	var body io.Reader
	if formData != nil {
		body = strings.NewReader(formData.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, method, link, body)
	if err != nil {
		return
	}

	if formData != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if queryParams != nil {
		req.URL.RawQuery = queryParams.Encode()
	}

	return a.makeRequest(req, v)
}

func (a *Api) doRequestWithMultiPartForm(ctx context.Context, method, path, header string, queryParams url.Values, body io.Reader, v any) (err error) {
	var link = fmt.Sprintf("%s%s", a.address, path)
	req, err := http.NewRequestWithContext(ctx, method, link, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", header)

	if queryParams != nil {
		req.URL.RawQuery = queryParams.Encode()
	}

	return a.makeRequest(req, v)
}

func (a *Api) makeRequest(req *http.Request, v any) (err error) {
	resp, err := a.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if a.debug {
		fmt.Printf("received response status code: %v\n", resp.StatusCode)
	}

	if resp.StatusCode != 200 {
		content, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			err = err1
			return
		}
		err = errors.New(string(content))
		return
	}

	var rs io.ReadCloser
	if a.debug {
		content, _ := io.ReadAll(resp.Body)
		fmt.Println("received response body")
		fmt.Println(string(content))
		rs = io.NopCloser(bytes.NewReader(content))
	} else {
		rs = resp.Body
	}

	switch v2 := v.(type) {
	case *string:
		var content []byte
		content, err = io.ReadAll(rs)
		if err != nil {
			return
		}
		*v2 = string(content)
	default:
		err = json.NewDecoder(rs).Decode(v)
		if err != nil {
			return
		}
	}
	return
}
