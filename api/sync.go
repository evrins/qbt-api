package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type Sync service

type MainDataResponse struct {
	Rid               int64               `json:"rid"`
	FullUpdate        bool                `json:"full_update"`
	Torrents          map[string]Torrent  `json:"torrents"`
	TorrentsRemoved   []string            `json:"torrents_removed"`
	Categories        map[string]Category `json:"categories"`
	CategoriesRemoved []string            `json:"categories_removed"`
	Tags              []string            `json:"tags"`
	TagsRemoved       []string            `json:"tags_removed"`
	ServerState       ServerState         `json:"server_state"`
	Trackers          map[string][]string `json:"trackers"`
}

type Category struct {
	Name     string `json:"name"`
	SavePath string `json:"savePath"`
}

type ServerState struct {
	AllTimeDL            int    `json:"alltime_dl"`
	AllTimeUL            int    `json:"alltime_ul"`
	AverageTimeQueue     int    `json:"average_time_queue"`
	ConnectionStatus     string `json:"connection_status"`
	DhtNodes             int    `json:"dht_nodes"`
	DlInfoData           int    `json:"dl_info_data"`
	DlInfoSpeed          int    `json:"dl_info_speed"`
	DlRateLimit          int    `json:"dl_rate_limit"`
	FreeSpaceOnDisk      int64  `json:"free_space_on_disk"`
	GlobalRatio          string `json:"global_ratio"`
	QueuedIoJobs         int    `json:"queued_io_jobs"`
	Queueing             bool   `json:"queueing"`
	ReadCacheHits        string `json:"read_cache_hits"`
	ReadCacheOverload    string `json:"read_cache_overload"`
	RefreshInterval      int    `json:"refresh_interval"`
	TotalBuffersSize     int    `json:"total_buffers_size"`
	TotalPeerConnections int    `json:"total_peer_connections"`
	TotalQueuedSize      int    `json:"total_queued_size"`
	TotalWastedSession   int    `json:"total_wasted_session"`
	UpInfoData           int    `json:"up_info_data"`
	UpInfoSpeed          int    `json:"up_info_speed"`
	UpRateLimit          int    `json:"up_rate_limit"`
	UseAltSpeedLimits    bool   `json:"use_alt_speed_limits"`
	WriteCacheOverload   string `json:"write_cache_overload"`
}

type Torrent struct {
	AddedOn           int    `json:"added_on"`
	AmountLeft        int64  `json:"amount_left"`
	AutoTmm           bool   `json:"auto_tmm"`
	Availability      int    `json:"availability"`
	Category          string `json:"category"`
	Completed         int    `json:"completed"`
	CompletionOn      int    `json:"completion_on"`
	ContentPath       string `json:"content_path"`
	DlLimit           int    `json:"dl_limit"`
	Dlspeed           int    `json:"dlspeed"`
	DownloadPath      string `json:"download_path"`
	Downloaded        int    `json:"downloaded"`
	DownloadedSession int    `json:"downloaded_session"`
	Eta               int    `json:"eta"`
	FLPiecePrio       bool   `json:"f_l_piece_prio"`
	ForceStart        bool   `json:"force_start"`
	InfohashV1        string `json:"infohash_v1"`
	InfohashV2        string `json:"infohash_v2"`
	LastActivity      int    `json:"last_activity"`
	MagnetURI         string `json:"magnet_uri"`
	MaxRatio          int    `json:"max_ratio"`
	MaxSeedingTime    int    `json:"max_seeding_time"`
	Name              string `json:"name"`
	NumComplete       int    `json:"num_complete"`
	NumIncomplete     int    `json:"num_incomplete"`
	NumLeechs         int    `json:"num_leechs"`
	NumSeeds          int    `json:"num_seeds"`
	Priority          int    `json:"priority"`
	Progress          int    `json:"progress"`
	Ratio             int    `json:"ratio"`
	RatioLimit        int    `json:"ratio_limit"`
	SavePath          string `json:"save_path"`
	SeedingTime       int    `json:"seeding_time"`
	SeedingTimeLimit  int    `json:"seeding_time_limit"`
	SeenComplete      int    `json:"seen_complete"`
	SeqDl             bool   `json:"seq_dl"`
	Size              int64  `json:"size"`
	State             string `json:"state"`
	SuperSeeding      bool   `json:"super_seeding"`
	Tags              string `json:"tags"`
	TimeActive        int    `json:"time_active"`
	TotalSize         int64  `json:"total_size"`
	Tracker           string `json:"tracker"`
	TrackersCount     int    `json:"trackers_count"`
	UpLimit           int    `json:"up_limit"`
	Uploaded          int    `json:"uploaded"`
	UploadedSession   int    `json:"uploaded_session"`
	Upspeed           int    `json:"upspeed"`
}

func (s *Sync) MainData(ctx context.Context, rid int64) (mainData MainDataResponse, err error) {
	path := "/api/v2/sync/maindata"
	query := url.Values{}
	query.Set("rid", strconv.FormatInt(rid, 10))

	resp, _, err := s.api.doRequest(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	err = json.NewDecoder(resp).Decode(&mainData)
	if err != nil {
		return
	}
	return
}

type TorrentPeersResponse struct {
	FullUpdate bool            `json:"full_update"`
	Peers      map[string]Peer `json:"peers"`
	Rid        int             `json:"rid"`
	ShowFlags  bool            `json:"show_flags"`
}

type Peer struct {
	Client       string `json:"client"`
	Connection   string `json:"connection"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	DlSpeed      int    `json:"dl_speed"`
	Downloaded   int    `json:"downloaded"`
	Files        string `json:"files"`
	Flags        string `json:"flags"`
	FlagsDesc    string `json:"flags_desc"`
	IP           string `json:"ip"`
	PeerIDClient string `json:"peer_id_client"`
	Port         int    `json:"port"`
	Progress     int    `json:"progress"`
	Relevance    int    `json:"relevance"`
	UpSpeed      int    `json:"up_speed"`
	Uploaded     int    `json:"uploaded"`
}

func (s *Sync) TorrentPeers(ctx context.Context, hash string, rid int64) (torrentPeersResponse TorrentPeersResponse, err error) {
	path := "/api/v2/sync/torrentPeers"

	query := url.Values{}
	query.Add("hash", hash)
	query.Add("rid", strconv.FormatInt(rid, 10))

	resp, _, err := s.api.doRequest(ctx, http.MethodGet, path, query, nil)

	if err != nil {
		return
	}
	defer resp.Close()
	err = json.NewDecoder(resp).Decode(&torrentPeersResponse)
	if err != nil {
		return
	}
	return
}
