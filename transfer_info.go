package qbt_api

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type TransferInfo service

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
	path := "/api/v2/transfer/info"

	err = ti.api.doRequest(ctx, http.MethodGet, path, nil, nil, &info)
	if err != nil {
		return
	}
	return
}

type SpeedLimitsMode string

const AlternativeSpeedLimitsDisabled SpeedLimitsMode = "0"
const AlternativeSpeedLimitsEnabled SpeedLimitsMode = "1"

func (ti *TransferInfo) SpeedLimitsMode(ctx context.Context) (speedLimitsMode SpeedLimitsMode, err error) {
	path := "/api/v2/transfer/speedLimitsMode"

	err = ti.api.doRequest(ctx, http.MethodGet, path, nil, nil, &speedLimitsMode)
	if err != nil {
		return
	}
	return
}

func (ti *TransferInfo) ToggleSpeedLimitsMode(ctx context.Context) (err error) {
	path := "/api/v2/transfer/toggleSpeedLimitsMode"

	err = ti.api.doRequest(ctx, http.MethodPost, path, nil, nil, emptyResponse)
	if err != nil {
		return
	}
	return
}

// DownloadLimit return current download global speed limit in bytes/second return zero if no limit
func (ti *TransferInfo) DownloadLimit(ctx context.Context) (limit int64, err error) {
	path := "/api/v2/transfer/downloadLimit"

	err = ti.api.doRequest(ctx, http.MethodGet, path, nil, nil, &limit)
	if err != nil {
		return
	}
	return
}

func (ti *TransferInfo) SetDownloadLimit(ctx context.Context, limit int64) (err error) {
	path := "/api/v2/transfer/setDownloadLimit"

	formData := url.Values{
		"limit": []string{strconv.FormatInt(limit, 10)},
	}

	err = ti.api.doRequest(ctx, http.MethodPost, path, nil, formData, emptyResponse)
	if err != nil {
		return
	}
	return
}

// UploadLimit return current upload global speed limit in bytes/second return zero if no limit
func (ti *TransferInfo) UploadLimit(ctx context.Context) (limit int64, err error) {
	path := "/api/v2/transfer/uploadLimit"

	err = ti.api.doRequest(ctx, http.MethodGet, path, nil, nil, &limit)
	if err != nil {
		return
	}
	return
}

func (ti *TransferInfo) SetUploadLimit(ctx context.Context, limit int64) (err error) {
	path := "/api/v2/transfer/setUploadLimit"

	formData := url.Values{
		"limit": []string{strconv.FormatInt(limit, 10)},
	}

	err = ti.api.doRequest(ctx, http.MethodPost, path, nil, formData, emptyResponse)
	if err != nil {
		return
	}

	return
}

func (ti *TransferInfo) BanPeers(ctx context.Context, peerList []string) (err error) {
	path := "/api/v2/transfer/banPeers"

	peers := strings.Join(peerList, "|")
	formData := url.Values{}
	formData.Set("peers", peers)

	err = ti.api.doRequest(ctx, http.MethodPost, path, nil, formData, emptyResponse)
	return
}
