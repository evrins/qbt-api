package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type Option func(api *Api)

func EnableDebug(api *Api) {
	api.debug = true
}

type Api struct {
	hc                *http.Client
	address           string
	debug             bool
	Auth              *Auth
	App               *App
	Log               *Log
	Sync              *Sync
	TransferInfo      *TransferInfo
	TorrentManagement *TorrentManagement
	Rss               *Rss
}

func NewApi(address string, options ...Option) (api *Api, err error) {
	api = &Api{}
	api.address = strings.TrimSuffix(address, "/")
	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}
	api.hc = &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	for _, opt := range options {
		opt(api)
	}

	api.Auth = &Auth{api}
	api.App = &App{api}
	api.Log = &Log{api}
	api.Sync = &Sync{api}
	api.TransferInfo = &TransferInfo{api}
	api.TorrentManagement = &TorrentManagement{api}
	api.Rss = &Rss{api}

	return api, nil
}

func (a *Api) doRequest(ctx context.Context, method, link string, queryParams, formData url.Values) (rs io.ReadCloser, statusCode int, err error) {
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

	return a.makeRequest(req)
}

func (a *Api) doRequestWithMultiPartForm(ctx context.Context, method, link, header string, queryParams url.Values, body io.Reader) (rs io.ReadCloser, statusCode int, err error) {

	req, err := http.NewRequestWithContext(ctx, method, link, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", header)

	if queryParams != nil {
		req.URL.RawQuery = queryParams.Encode()
	}

	return a.makeRequest(req)
}

func (a *Api) makeRequest(req *http.Request) (rs io.ReadCloser, statusCode int, err error) {
	resp, err := a.hc.Do(req)
	if err != nil {
		return
	}

	if a.debug {
		fmt.Printf("received response status code: %v\n", resp.StatusCode)
		content, _ := io.ReadAll(resp.Body)
		fmt.Println("received response body")
		fmt.Println(string(content))
		rs = io.NopCloser(bytes.NewReader(content))
	} else {
		rs = resp.Body
	}

	statusCode = resp.StatusCode
	return
}
