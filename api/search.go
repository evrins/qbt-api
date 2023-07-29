package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Search struct {
	*Api
}

type SearchOptions struct {
	Pattern           string
	Plugins           []string
	UseAllPlugins     bool
	UseEnabledPlugins bool
	Category          []string
	UseAllCategory    bool
}

type StartResponse struct {
	Id int64 `json:"id"`
}

func (s *Search) Start(ctx context.Context, opts *SearchOptions) (startResponse *StartResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/search/start", s.address)

	formData := url.Values{}
	formData.Set("pattern", opts.Pattern)
	if opts.UseAllPlugins {
		formData.Set("plugins", "all")
	} else if opts.UseEnabledPlugins {
		formData.Set("plugins", "enabled")
	} else {
		formData.Set("plugins", strings.Join(opts.Plugins, "|"))
	}

	formData.Set("category", joinHashes(opts.Category, opts.UseAllCategory))

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	startResponse = &StartResponse{}
	err = json.NewDecoder(resp).Decode(&startResponse)
	return
}

func (s *Search) Stop(ctx context.Context, id int64) (err error) {
	link := fmt.Sprintf("%s/api/v2/search/stop", s.address)

	formData := url.Values{}
	formData.Set("id", strconv.FormatInt(id, 10))

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type StatusResponse []Status

type StatusEnum string

const StatusRunning StatusEnum = "Running"
const StatusStopped StatusEnum = "Stopped"

type Status struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
	Total  int64  `json:"total"`
}

func (s *Search) Status(ctx context.Context, id int64) (statusResponse StatusResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/search/status", s.address)

	query := url.Values{}
	if id != 0 {
		query.Set("id", strconv.FormatInt(id, 10))
	}

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	statusResponse = []Status{}
	err = json.NewDecoder(resp).Decode(&statusResponse)
	if err != nil {
		return
	}
	return
}

type ResultResponse struct {
	Results []Result   `json:"results"`
	Status  StatusEnum `json:"status"`
	Total   int64      `json:"total"`
}

type Result struct {
	DescrLink  string `json:"descrLink"`
	FileName   string `json:"fileName"`
	FileSize   int64  `json:"fileSize"`
	FileUrl    string `json:"fileUrl"`
	NbLeechers int64  `json:"nbLeechers"`
	NbSeeders  int64  `json:"nbSeeders"`
	SiteUrl    string `json:"siteUrl"`
}

func (s *Search) Results(ctx context.Context, id, limit, offset int64) (resultResponse *ResultResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/search/results", s.address)

	formData := url.Values{}
	formData.Set("id", strconv.FormatInt(id, 10))
	formData.Set("limit", strconv.FormatInt(limit, 10))
	formData.Set("offset", strconv.FormatInt(offset, 10))

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	resultResponse = &ResultResponse{}
	err = json.NewDecoder(resp).Decode(&resultResponse)
	if err != nil {
		return
	}
	return
}

func (s *Search) Delete(ctx context.Context, id int64) (err error) {
	link := fmt.Sprintf("%s/api/v2/search/delete", s.address)

	formData := url.Values{}
	if id != 0 {
		formData.Set("id", strconv.FormatInt(id, 10))
	}

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type PluginsResponse []*Plugin

type Plugin struct {
	Enabled             bool             `json:"enabled"`
	FullName            string           `json:"fullName"`
	Name                string           `json:"name"`
	SupportedCategories []PluginCategory `json:"supportedCategories"`
	Url                 string           `json:"url"`
	Version             string           `json:"version"`
}

type PluginCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (s *Search) Plugins(ctx context.Context) (pluginsResponse PluginsResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/search/plugins", s.address)

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	pluginsResponse = []*Plugin{}
	err = json.NewDecoder(resp).Decode(&pluginsResponse)
	if err != nil {
		return
	}
	return
}

func (s *Search) InstallPlugin(ctx context.Context, sources []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/search/installPlugin", s.address)

	formData := url.Values{}
	formData.Set("sources", strings.Join(sources, "|"))

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (s *Search) UninstallPlugin(ctx context.Context, names []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/search/uninstallPlugin", s.address)

	formData := url.Values{}
	formData.Set("names", strings.Join(names, "|"))

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (s *Search) EnablePlugin(ctx context.Context, names []string, enable bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/search/enablePlugin", s.address)

	formData := url.Values{}
	formData.Set("names", strings.Join(names, "|"))
	formData.Set("enable", strconv.FormatBool(enable))

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (s *Search) UpdatePlugins(ctx context.Context) (err error) {
	link := fmt.Sprintf("%s/api/v2/search/updatePlugins", s.address)

	resp, _, err := s.doRequest(ctx, http.MethodPost, link, nil, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}
