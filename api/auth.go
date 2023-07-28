package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Auth struct {
	*Api
}

func (a *Auth) Login(ctx context.Context, username, password string) (respText string, err error) {
	link := fmt.Sprintf("%s/api/v2/auth/login", a.address)

	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	body := strings.NewReader(formData.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	respText = string(content)
	return
}

func (a *Auth) Logout(ctx context.Context) (respText string, err error) {
	link := fmt.Sprintf("%s/api/v2/auth/logout", a.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, nil)
	if err != nil {
		return
	}
	resp, err := a.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	respText = string(content)
	return
}
