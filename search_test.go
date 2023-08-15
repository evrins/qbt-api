package qbt_api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

func TestSearch_Start(t *testing.T) {
	opts := &SearchOptions{
		Pattern:           "life is like a boat",
		Plugins:           nil,
		UseAllPlugins:     true,
		UseEnabledPlugins: false,
		Category:          nil,
		UseAllCategory:    true,
	}

	var resp, err = api.Search.Start(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(resp)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var counter = 0

	for range ticker.C {
		var status StatusResponse
		status, err = api.Search.Status(context.Background(), resp.Id)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(status)

		var results *ResultResponse
		results, err = api.Search.Results(context.Background(), resp.Id, 0, 0)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(results)

		counter += 1
		if counter > 2 {
			break
		}
	}

	err = api.Search.Delete(context.Background(), resp.Id)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	var status StatusResponse
	status, err = api.Search.Status(context.Background(), resp.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(status)
}

func TestSearch_Plugins(t *testing.T) {
	var resp, err = api.Search.Plugins(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(resp)
}

func TestSearch_EnablePlugin(t *testing.T) {
	var names = []string{"kickass_torrent"}
	var enable = true
	var err = api.Search.EnablePlugin(context.Background(), names, enable)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearch_UpdatePlugins(t *testing.T) {
	var err = api.Search.UpdatePlugins(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
