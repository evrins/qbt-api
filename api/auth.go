package api

import (
	"context"
	"net/http"
	"net/url"
)

type Auth service

func (a *Auth) Login(ctx context.Context, username, password string) (respText string, err error) {
	path := "/api/v2/auth/login"

	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	err = a.api.doRequest(ctx, http.MethodPost, path, nil, formData, &respText)
	if err != nil {
		return
	}
	return
}

func (a *Auth) Logout(ctx context.Context) (respText string, err error) {
	path := "/api/v2/auth/logout"

	err = a.api.doRequest(ctx, http.MethodPost, path, nil, nil, &respText)

	if err != nil {
		return
	}
	return
}
