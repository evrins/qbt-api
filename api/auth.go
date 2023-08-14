package api

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type Auth service

func (a *Auth) Login(ctx context.Context, username, password string) (respText string, err error) {
	path := "/api/v2/auth/login"

	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	resp, _, err := a.api.doRequest(ctx, http.MethodPost, path, nil, formData)
	if err != nil {
		return
	}

	defer resp.Close()
	content, err := io.ReadAll(resp)
	if err != nil {
		return
	}
	respText = string(content)
	return
}

func (a *Auth) Logout(ctx context.Context) (respText string, err error) {
	path := "/api/v2/auth/logout"

	resp, _, err := a.api.doRequest(ctx, http.MethodPost, path, nil, nil)

	if err != nil {
		return
	}
	defer resp.Close()
	content, err := io.ReadAll(resp)
	if err != nil {
		return
	}
	respText = string(content)
	return
}
