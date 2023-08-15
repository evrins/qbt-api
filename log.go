package qbt_api

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

type Log service

type LogOptions struct {
	Normal      bool `json:"normal"`
	Info        bool `json:"info"`
	Warning     bool `json:"warning"`
	Critical    bool `json:"critical"`
	LastKnownId int  `json:"last_known_id"`
}

var DefaultLogOptions = LogOptions{
	Normal:      true,
	Info:        true,
	Warning:     true,
	Critical:    true,
	LastKnownId: -1,
}

type LogType int

const Normal LogType = 1
const Info LogType = 2
const Warning LogType = 4
const Critical LogType = 8

type LogItem struct {
	Id        int64   `json:"id"`
	Message   string  `json:"message"`
	Timestamp int64   `json:"timestamp"`
	Type      LogType `json:"type"`
}

// Main copy and modify DefaultLogOptions and pass modified options as parameter for default value is true
func (l *Log) Main(ctx context.Context, opts LogOptions) (logs []*LogItem, err error) {
	path := "/api/v2/log/main"

	query := url.Values{}
	query.Set("normal", strconv.FormatBool(opts.Normal))
	query.Set("info", strconv.FormatBool(opts.Info))
	query.Set("warning", strconv.FormatBool(opts.Warning))
	query.Set("critical", strconv.FormatBool(opts.Critical))
	query.Set("last_known_id", strconv.Itoa(opts.LastKnownId))

	err = l.api.doRequest(ctx, http.MethodGet, path, query, nil, &logs)
	if err != nil {
		return
	}

	return
}

type PeerLogOptions struct {
	LastKnownId int `json:"last_known_id"`
}

var DefaultPeerLogOptions = PeerLogOptions{LastKnownId: -1}

type PeerLogItem struct {
	Id        int64  `json:"id"`
	IP        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
	Blocked   bool   `json:"blocked"`
	Reason    string `json:"reason"`
}

// Peers copy and modify DefaultPeerLogOptions and pass modified options as parameter for default value is -1
func (l *Log) Peers(ctx context.Context, opts PeerLogOptions) (logs []*PeerLogItem, err error) {
	path := "/api/v2/log/peers"

	query := url.Values{}
	query.Set("last_known_id", strconv.Itoa(opts.LastKnownId))

	err = l.api.doRequest(ctx, http.MethodGet, path, nil, nil, &logs)
	if err != nil {
		return
	}

	return
}
