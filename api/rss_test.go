package api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestRss_AddFolder(t *testing.T) {
	var folder = "f1"

	var err = api.Rss.AddFolder(context.Background(), folder)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRss_AddFeed(t *testing.T) {
	var url_ = "https://dmhy.org/topics/rss/rss.xml"
	var path = "rss1"

	var err = api.Rss.AddFeed(context.Background(), url_, path)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRss_RemoveItem_Folder(t *testing.T) {
	var folder = "f1"

	var err = api.Rss.RemoveItem(context.Background(), folder)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRss_RemoveItem_Feed(t *testing.T) {
	var path = "f1/rss1"

	var err = api.Rss.RemoveItem(context.Background(), path)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRss_MoveItem(t *testing.T) {
	var itemPath = "rss1"
	var destPath = "f1/rss1"

	var err = api.Rss.MoveItem(context.Background(), itemPath, destPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRss_Items(t *testing.T) {
	var withData = true

	var resp, err = api.Rss.Items(context.Background(), withData)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(resp)
}

func TestRss_MarkAsRead(t *testing.T) {
	var itemPath = "f1/rss1"
	var articleId = ""

	var err = api.Rss.MarkAsRead(context.Background(), itemPath, articleId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRss_RefreshItem(t *testing.T) {
	var itemPath = "rss1"

	var err = api.Rss.RefreshItem(context.Background(), itemPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRss_SetRule(t *testing.T) {
	var ruleName = "ani"
	var ruleDef = &RuleDef{
		Enable:                    true,
		MustContain:               "BLEACH 死神 千年血戰篇-訣別譚",
		MustNotContain:            "",
		UseRegex:                  false,
		EpisodeFilter:             "",
		SmartFilter:               false,
		PreviouslyMatchedEpisodes: nil,
		AffectedFeeds: []string{
			"https://dmhy.org/topics/rss/rss.xml",
		},
		IgnoreDays:       0,
		LastMatch:        "",
		AddPaused:        false,
		AssignedCategory: "",
		SavePath:         "",
	}

	var err = api.Rss.SetRule(context.Background(), ruleName, ruleDef)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRss_RenameRule(t *testing.T) {
	var ruleName = "ani"
	var newRuleName = "ani0"

	var err = api.Rss.RenameRule(context.Background(), ruleName, newRuleName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRss_RemoveRule(t *testing.T) {
	var ruleName = "ani0"
	var err = api.Rss.RemoveRule(context.Background(), ruleName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRss_Rules(t *testing.T) {
	var resp, err = api.Rss.Rules(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(resp)
}

func TestRss_MatchingArticles(t *testing.T) {
	var ruleName = "ani0"
	var resp, err = api.Rss.MatchingArticles(context.Background(), ruleName)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(resp)
}
