package api

import (
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type Api struct {
	hc      *http.Client
	address string
	Auth    *Auth
	App     *App
	Log     *Log
	Sync    *Sync
}

func NewApi(address string) (api *Api, err error) {
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

	api.Auth = &Auth{api}
	api.App = &App{api}
	api.Log = &Log{api}
	api.Sync = &Sync{api}

	return api, nil
}
