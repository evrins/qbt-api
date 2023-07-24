package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type TransferInfo struct {
	*Api
}

type ConnectionStatus string

const Connected ConnectionStatus = "connected"
const Firewalled ConnectionStatus = "firewalled"
const Disconnected ConnectionStatus = "disconnected"

type InfoResponse struct {
	ConnectionStatus ConnectionStatus `json:"connection_status"`
	DhtNodes         int              `json:"dht_nodes"`
	DlInfoData       int              `json:"dl_info_data"`
	DlInfoSpeed      int              `json:"dl_info_speed"`
	DlRateLimit      int              `json:"dl_rate_limit"`
	UpInfoData       int              `json:"up_info_data"`
	UpInfoSpeed      int              `json:"up_info_speed"`
	UpRateLimit      int              `json:"up_rate_limit"`
}

func (ti *TransferInfo) Info(ctx context.Context) (info *InfoResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/info", ti.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}
	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return
	}
	return
}

type SpeedLimitsMode string

const AlternativeSpeedLimitsDisabled SpeedLimitsMode = "0"
const AlternativeSpeedLimitsEnabled SpeedLimitsMode = "1"

func (ti *TransferInfo) SpeedLimitsMode(ctx context.Context) (speedLimitsMode SpeedLimitsMode, err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/speedLimitsMode", ti.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}
	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	speedLimitsMode = SpeedLimitsMode(content)
	return
}

func (ti *TransferInfo) ToggleSpeedLimitsMode(ctx context.Context) (err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/toggleSpeedLimitsMode", ti.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, nil)
	if err != nil {
		return
	}
	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

// DownloadLimit return current download global speed limit in bytes/second return zero if no limit
func (ti *TransferInfo) DownloadLimit(ctx context.Context) (limit int64, err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/downloadLimit", ti.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}
	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	limit, err = strconv.ParseInt(string(content), 10, 64)
	if err != nil {
		return
	}
	return
}

func (ti *TransferInfo) SetDownloadLimit(ctx context.Context, limit int64) (err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/setDownloadLimit", ti.address)

	formData := url.Values{
		"limit": []string{strconv.FormatInt(limit, 10)},
	}
	body := strings.NewReader(formData.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

// UploadLimit return current upload global speed limit in bytes/second return zero if no limit
func (ti *TransferInfo) UploadLimit(ctx context.Context) (limit int64, err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/uploadLimit", ti.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}
	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	limit, err = strconv.ParseInt(string(content), 10, 64)
	if err != nil {
		return
	}
	return
}

func (ti *TransferInfo) SetUploadLimit(ctx context.Context, limit int64) (err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/setUploadLimit", ti.address)

	formData := url.Values{
		"limit": []string{strconv.FormatInt(limit, 10)},
	}
	body := strings.NewReader(formData.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func (ti *TransferInfo) BanPeers(ctx context.Context, peerList []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/transfer/banPeers", ti.address)

	peers := strings.Join(peerList, "|")
	formData := url.Values{}
	formData.Set("peers", peers)
	body := strings.NewReader(formData.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, body)
	if err != nil {
		return
	}

	resp, err := ti.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}
