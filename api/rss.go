package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type Rss service

func (r *Rss) AddFolder(ctx context.Context, path string) (err error) {
	path_ := "/api/v2/rss/addFolder"

	formData := url.Values{}
	formData.Set("path", path)

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path_, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (r *Rss) AddFeed(ctx context.Context, url_, path string) (err error) {
	path_ := "/api/v2/rss/addFeed"

	formData := url.Values{}
	formData.Set("url", url_)
	if path != "" {
		formData.Set("path", path)
	}

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path_, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (r *Rss) RemoveItem(ctx context.Context, path string) (err error) {
	path_ := "/api/v2/rss/removeItem"

	formData := url.Values{}
	formData.Set("path", path)

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path_, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (r *Rss) MoveItem(ctx context.Context, itemPath, destPath string) (err error) {
	path := "/api/v2/rss/moveItem"

	formData := url.Values{}
	formData.Set("itemPath", itemPath)
	formData.Set("destPath", destPath)

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type RssItemResponse map[string]RssItem

type RssItem struct {
	Articles      []RssArticle `json:"articles"`
	HasError      bool         `json:"hasError"`
	IsLoading     bool         `json:"isLoading"`
	LastBuildDate string       `json:"lastBuildDate"`
	Title         string       `json:"title"`
	UID           string       `json:"uid"`
	URL           string       `json:"url"`
}

type RssArticle struct {
	Author      string `json:"author"`
	Category    string `json:"category"`
	Date        string `json:"date"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Link        string `json:"link"`
	Title       string `json:"title"`
	TorrentURL  string `json:"torrentURL"`
}

func (r *Rss) Items(ctx context.Context, withData bool) (rssItemResponse RssItemResponse, err error) {
	path := "/api/v2/rss/items"

	query := url.Values{}
	query.Set("withData", strconv.FormatBool(withData))

	resp, _, err := r.api.doRequest(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	rssItemResponse = map[string]RssItem{}
	err = json.NewDecoder(resp).Decode(&rssItemResponse)
	if err != nil {
		return
	}
	return
}

func (r *Rss) MarkAsRead(ctx context.Context, itemPath, articleId string) (err error) {
	path := "/api/v2/rss/markAsRead"

	formData := url.Values{}
	formData.Set("itemPath", itemPath)
	if articleId != "" {
		formData.Set("articleId", articleId)
	}

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (r *Rss) RefreshItem(ctx context.Context, itemPath string) (err error) {
	path := "/api/v2/rss/refreshItem"

	formData := url.Values{}
	formData.Set("itemPath", itemPath)

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type RuleDef struct {
	Enable                    bool     `json:"enable"`
	MustContain               string   `json:"mustContain"`
	MustNotContain            string   `json:"mustNotContain"`
	UseRegex                  bool     `json:"useRegex"`
	EpisodeFilter             string   `json:"episodeFilter"`
	SmartFilter               bool     `json:"smartFilter"`
	PreviouslyMatchedEpisodes []string `json:"previouslyMatchedEpisodes"`
	AffectedFeeds             []string `json:"affectedFeeds"`
	IgnoreDays                int64    `json:"ignoreDays"`
	LastMatch                 string   `json:"lastMatch"`
	AddPaused                 bool     `json:"addPaused"`
	AssignedCategory          string   `json:"assignedCategory"`
	SavePath                  string   `json:"savePath"`
}

func (r *Rss) SetRule(ctx context.Context, ruleName string, ruleDef *RuleDef) (err error) {
	path := "/api/v2/rss/setRule"

	content, err := json.Marshal(ruleDef)
	if err != nil {
		return
	}

	formData := url.Values{}
	formData.Set("ruleName", ruleName)
	formData.Set("ruleDef", string(content))

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (r *Rss) RenameRule(ctx context.Context, ruleName, newRuleName string) (err error) {
	path := "/api/v2/rss/renameRule"

	formData := url.Values{}
	formData.Set("ruleName", ruleName)
	formData.Set("newRuleName", newRuleName)

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (r *Rss) RemoveRule(ctx context.Context, ruleName string) (err error) {
	path := "/api/v2/rss/removeRule"

	formData := url.Values{}
	formData.Set("ruleName", ruleName)

	resp, _, err := r.api.doRequest(ctx, http.MethodPost, path, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type RulesResponse map[string]RuleDef

func (r *Rss) Rules(ctx context.Context) (rulesResponse RulesResponse, err error) {
	path := "/api/v2/rss/rules"

	resp, _, err := r.api.doRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	rulesResponse = map[string]RuleDef{}
	err = json.NewDecoder(resp).Decode(&rulesResponse)
	if err != nil {
		return
	}
	return
}

type MatchingArticleResponse map[string][]string

func (r *Rss) MatchingArticles(ctx context.Context, ruleName string) (matchingArticleResponse MatchingArticleResponse, err error) {
	path := "/api/v2/rss/matchingArticles"

	query := url.Values{}
	query.Set("ruleName", ruleName)

	resp, _, err := r.api.doRequest(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	matchingArticleResponse = map[string][]string{}
	err = json.NewDecoder(resp).Decode(&matchingArticleResponse)
	if err != nil {
		return
	}
	return
}
