package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type App service

func (a *App) Version(ctx context.Context) (respText string, err error) {
	var path = "/api/v2/app/version"
	err = a.api.doRequest(ctx, http.MethodPost, path, nil, nil, &respText)
	if err != nil {
		return
	}
	return
}

func (a *App) WebApiVersion(ctx context.Context) (respText string, err error) {
	path := "/api/v2/app/webapiVersion"
	err = a.api.doRequest(ctx, http.MethodPost, path, nil, nil, &respText)
	if err != nil {
		return
	}
	return
}

type BuildInfo struct {
	QT         string `json:"qt"`
	Libtorrent string `json:"libtorrent"`
	Boost      string `json:"boost"`
	Openssl    string `json:"openssl"`
	Bitness    int    `json:"bitness"`
}

func (a *App) BuildInfo(ctx context.Context) (bi *BuildInfo, err error) {
	path := "/api/v2/app/buildInfo"
	err = a.api.doRequest(ctx, http.MethodPost, path, nil, nil, &bi)
	if err != nil {
		return
	}
	return
}

func (a *App) Shutdown(ctx context.Context) (respText string, err error) {
	path := "/api/v2/app/shutdown"
	err = a.api.doRequest(ctx, http.MethodPost, path, nil, nil, &respText)
	if err != nil {
		return
	}
	return
}

type SchedulerDays int

const EveryDay SchedulerDays = 0
const EveryWeekday SchedulerDays = 1
const EveryWeekend SchedulerDays = 2
const EveryMonday SchedulerDays = 3
const EveryTuesday SchedulerDays = 4
const EveryWednesday SchedulerDays = 5
const EveryThursday SchedulerDays = 6
const EveryFriday SchedulerDays = 7
const EverySaturday SchedulerDays = 8
const EverySunday SchedulerDays = 9

type Encryption int

const PreferEncryption Encryption = 0
const ForceEncryptionOn Encryption = 1
const ForceEncryptionOff Encryption = 2

type ProxyType int

const ProxyDisabled ProxyType = -1
const ProxyWithoutAuthentication ProxyType = 1
const Socks5ProxyWithoutAuthentication ProxyType = 2
const HttpProxyWithAuthentication ProxyType = 3
const Socks5ProxyWithAuthentication ProxyType = 4
const Socks4ProxyWithoutAuthentication ProxyType = 5

type DyndnsService int

const UseDyDNS DyndnsService = 0
const UseNoIP DyndnsService = 1

type MaxRatioAct int

const PauseTorrent MaxRatioAct = 0
const RemoveTorrent MaxRatioAct = 1

type BittorrentProtocol int

const TCPAndUTP BittorrentProtocol = 0
const TCP BittorrentProtocol = 1
const UTP BittorrentProtocol = 2

type UploadChokingAlgorithm int

const RoundRobin UploadChokingAlgorithm = 0
const FastestUpload UploadChokingAlgorithm = 1
const AntiLeech UploadChokingAlgorithm = 2

type UploadSlotsBehavior int

const FixedSlots UploadSlotsBehavior = 0
const UploadRateBased UploadSlotsBehavior = 1

type UTPTCPMixedMode int

const PreferTCP UTPTCPMixedMode = 0
const PeerProportional UTPTCPMixedMode = 1

type Preferences struct {
	Locale                             string                 `json:"locale"`
	CreateSubfolderEnabled             bool                   `json:"create_subfolder_enabled"`
	StartPausedEnabled                 bool                   `json:"start_paused_enabled"`
	AutoDeleteMode                     int                    `json:"auto_delete_mode"`
	PreallocateAll                     bool                   `json:"preallocate_all"`
	IncompleteFilesExt                 bool                   `json:"incomplete_files_ext"`
	AutoTmmEnabled                     bool                   `json:"auto_tmm_enabled"`
	TorrentChangedTmmEnabled           bool                   `json:"torrent_changed_tmm_enabled"`
	SavePathChangedTmmEnabled          bool                   `json:"save_path_changed_tmm_enabled"`
	CategoryChangedTmmEnabled          bool                   `json:"category_changed_tmm_enabled"`
	SavePath                           string                 `json:"save_path"`
	TempPathEnabled                    bool                   `json:"temp_path_enabled"`
	TempPath                           string                 `json:"temp_path"`
	ScanDirs                           map[string]any         `json:"scan_dirs"` // 0 Download to the monitored folder, 1 Download to the default save path, string Download to this path
	ExportDir                          string                 `json:"export_dir"`
	ExportDirFin                       string                 `json:"export_dir_fin"`
	MailNotificationEnabled            bool                   `json:"mail_notification_enabled"`
	MailNotificationSender             string                 `json:"mail_notification_sender"`
	MailNotificationEmail              string                 `json:"mail_notification_email"`
	MailNotificationSmtp               string                 `json:"mail_notification_smtp"`
	MailNotificationSSLEnabled         bool                   `json:"mail_notification_ssl_enabled"`
	MailNotificationAuthEnabled        bool                   `json:"mail_notification_auth_enabled"`
	MailNotificationUsername           string                 `json:"mail_notification_username"`
	MailNotificationPassword           string                 `json:"mail_notification_password"`
	AutorunEnabled                     bool                   `json:"autorun_enabled"`
	AutorunProgram                     string                 `json:"autorun_program"`
	QueueingEnabled                    bool                   `json:"queueing_enabled"`
	MaxActiveDownloads                 int64                  `json:"max_active_downloads"`
	MaxActiveTorrents                  int64                  `json:"max_active_torrents"`
	MaxActiveUploads                   int64                  `json:"max_active_uploads"`
	DontCountSlowTorrents              bool                   `json:"dont_count_slow_torrents"`
	SlowTorrentDLRateThreshold         int64                  `json:"slow_torrent_dl_rate_threshold"`
	SlowTorrentULRateThreshold         int64                  `json:"slow_torrent_ul_rate_threshold"`
	SlowTorrentInactiveTimer           int64                  `json:"slow_torrent_inactive_timer"`
	MaxRatioEnabled                    bool                   `json:"max_ratio_enabled"`
	MaxRatio                           float64                `json:"max_ratio"`
	MaxRatioAct                        MaxRatioAct            `json:"max_ratio_act"`
	ListenPort                         int64                  `json:"listen_port"`
	Upnp                               bool                   `json:"upnp"`
	RandomPort                         bool                   `json:"random_port"`
	DLLimit                            int64                  `json:"dl_limit"`
	UPLimit                            int64                  `json:"up_limit"`
	MaxConnec                          int64                  `json:"max_connec"`
	MaxConnecPerTorrent                int64                  `json:"max_connec_per_torrent"`
	MaxUploads                         int64                  `json:"max_uploads"`
	MaxUploadsPerTorrent               int64                  `json:"max_uploads_per_torrent"`
	StopTrackerTimeout                 int64                  `json:"stop_tracker_timeout"`
	EnablePieceExtentAffinity          bool                   `json:"enable_piece_extent_affinity"`
	BittorrentProtocol                 BittorrentProtocol     `json:"bittorrent_protocol"`
	LimitUTPRate                       bool                   `json:"limit_utp_rate"`
	LimitTCPOverhead                   bool                   `json:"limit_tcp_overhead"`
	LimitLanPeers                      bool                   `json:"limit_lan_peers"`
	AltDLLimit                         int64                  `json:"alt_dl_limit"`
	AltUPLimit                         int64                  `json:"alt_up_limit"`
	SchedulerEnabled                   bool                   `json:"scheduler_enabled"`
	ScheduleFromHour                   int64                  `json:"schedule_from_hour"`
	ScheduleFromMin                    int64                  `json:"schedule_from_min"`
	ScheduleToHour                     int64                  `json:"schedule_to_hour"`
	ScheduleToMin                      int64                  `json:"schedule_to_min"`
	SchedulerDays                      SchedulerDays          `json:"scheduler_days"`
	DHT                                bool                   `json:"dht"`
	PEX                                bool                   `json:"pex"`
	LSD                                bool                   `json:"lsd"`
	Encryption                         Encryption             `json:"encryption"`
	AnonymousMode                      bool                   `json:"anonymous_mode"`
	ProxyType                          ProxyType              `json:"proxy_type"`
	ProxyIP                            string                 `json:"proxy_ip"`
	ProxyPort                          int64                  `json:"proxy_port"`
	ProxyPeerConnections               bool                   `json:"proxy_peer_connections"`
	ProxyAuthEnabled                   bool                   `json:"proxy_auth_enabled"`
	ProxyUsername                      string                 `json:"proxy_username"`
	ProxyPassword                      string                 `json:"proxy_password"`
	ProxyTorrentsOnly                  bool                   `json:"proxy_torrents_only"`
	IPFilterEnabled                    bool                   `json:"ip_filter_enabled"`
	IPFilterPath                       string                 `json:"ip_filter_path"`
	IPFilterTrackers                   bool                   `json:"ip_filter_trackers"`
	WebUIDomainList                    string                 `json:"web_ui_domain_list"`
	WebUIAddress                       string                 `json:"web_ui_address"`
	WebUIPort                          int64                  `json:"web_ui_port"`
	WebUIUpnp                          bool                   `json:"web_ui_upnp"`
	WebUIUsername                      string                 `json:"web_ui_username"`
	WebUIPassword                      string                 `json:"web_ui_password"`
	WebUICSRFProtectionEnabled         bool                   `json:"web_ui_csrf_protection_enabled"`
	WebUIClickJackingProtectionEnabled bool                   `json:"web_ui_clickjacking_protection_enabled"`
	WebUISecureCookieEnabled           bool                   `json:"web_ui_secure_cookie_enabled"`
	WebUIMaxAuthFailCount              int64                  `json:"web_ui_max_auth_fail_count"`
	WebUIBanDuration                   int64                  `json:"web_ui_ban_duration"`
	WebUISessionTimeout                int64                  `json:"web_ui_session_timeout"`
	WebUIHostHeadeValidationEnabled    bool                   `json:"web_ui_host_header_validation_enabled"`
	BypassLocalAuth                    bool                   `json:"bypass_local_auth"`
	BypassAuthSubnetWhitelistEnabled   bool                   `json:"bypass_auth_subnet_whitelist_enabled"`
	BypassAuthSubnetWhitelist          string                 `json:"bypass_auth_subnet_whitelist"`
	AlternativeWebUIEnabled            bool                   `json:"alternative_webui_enabled"`
	AlternativeEebUIPath               string                 `json:"alternative_webui_path"`
	UseHttps                           bool                   `json:"use_https"`
	SSLKey                             string                 `json:"ssl_key"`
	SSLCert                            string                 `json:"ssl_cert"`
	WebUIHttpsKeyPath                  string                 `json:"web_ui_https_key_path"`
	WebUIHttpsCertPath                 string                 `json:"web_ui_https_cert_path"`
	DyndnsEnabled                      bool                   `json:"dyndns_enabled"`
	DyndnsService                      DyndnsService          `json:"dyndns_service"`
	DyndnsUsername                     string                 `json:"dyndns_username"`
	DyndnsPassword                     string                 `json:"dyndns_password"`
	DyndnsDomain                       string                 `json:"dyndns_domain"`
	RssRefreshInterval                 int64                  `json:"rss_refresh_interval"`
	RssMaxArticlesPerFeed              int64                  `json:"rss_max_articles_per_feed"`
	RssProcessingEnabled               bool                   `json:"rss_processing_enabled"`
	RssAutoDownloadingEnabled          bool                   `json:"rss_auto_downloading_enabled"`
	RssDownloadRepackProperEpisodes    bool                   `json:"rss_download_repack_proper_episodes"`
	RssSmartEpisodeFilters             string                 `json:"rss_smart_episode_filters"`
	AddTrackersEnabled                 bool                   `json:"add_trackers_enabled"`
	AddTrackers                        string                 `json:"add_trackers"`
	WebUIUseCustomHttpHeadersEnabled   bool                   `json:"web_ui_use_custom_http_headers_enabled"`
	WebUICustomHttpHeaders             string                 `json:"web_ui_custom_http_headers"`
	MaxSeedingTimeEnabled              bool                   `json:"max_seeding_time_enabled"`
	MaxSeedingTime                     int64                  `json:"max_seeding_time"`
	AnnounceIP                         string                 `json:"announce_ip"`
	AnnounceToAllTiers                 bool                   `json:"announce_to_all_tiers"`
	AnnounceToAllTrackers              bool                   `json:"announce_to_all_trackers"`
	AsyncIOThreads                     int64                  `json:"async_io_threads"`
	BannedIPs                          string                 `json:"banned_IPs"`
	CheckingMemoryUse                  int64                  `json:"checking_memory_use"`
	CurrentInterfaceAddress            string                 `json:"current_interface_address"`
	CurrentNetworkInterface            string                 `json:"current_network_interface"`
	DiskCache                          int64                  `json:"disk_cache"`
	DiskCacheTTL                       int64                  `json:"disk_cache_ttl"`
	EmbeddedTrackerPort                int64                  `json:"embedded_tracker_port"`
	EnableCoalesceReadWrite            bool                   `json:"enable_coalesce_read_write"`
	EnableEmbeddedTracker              bool                   `json:"enable_embedded_tracker"`
	EnableMultiConnectionsFromSameIP   bool                   `json:"enable_multi_connections_from_same_ip"`
	EnableOSCache                      bool                   `json:"enable_os_cache"`
	EnableUploadSuggestions            bool                   `json:"enable_upload_suggestions"`
	FilePoolSize                       int64                  `json:"file_pool_size"`
	OutgoingPortsMax                   int64                  `json:"outgoing_ports_max"`
	OutgoingPortsMin                   int64                  `json:"outgoing_ports_min"`
	RecheckCompletedTorrents           bool                   `json:"recheck_completed_torrents"`
	ResolvePeerCountries               bool                   `json:"resolve_peer_countries"`
	SaveResumeDataInterval             int64                  `json:"save_resume_data_interval"`
	SendBufferLowWatermark             int64                  `json:"send_buffer_low_watermark"`
	SendBufferWatermark                int64                  `json:"send_buffer_watermark"`
	SendBufferWatermarkFactor          int64                  `json:"send_buffer_watermark_factor"`
	SocketBacklogSize                  int64                  `json:"socket_backlog_size"`
	UploadChokingAlgorithm             UploadChokingAlgorithm `json:"upload_choking_algorithm"`
	UploadSlotsBehavior                UploadSlotsBehavior    `json:"upload_slots_behavior"`
	UpnpLeaseDuration                  int64                  `json:"upnp_lease_duration"`
	UTPTCPMixedMode                    UTPTCPMixedMode        `json:"utp_tcp_mixed_mode"`
}

func (a *App) Preferences(ctx context.Context) (preferences *Preferences, err error) {
	path := "/api/v2/app/preferences"
	err = a.api.doRequest(ctx, http.MethodPost, path, nil, nil, &preferences)
	if err != nil {
		return
	}
	return
}

func (a *App) SetPreferences(ctx context.Context, pref Preferences) (respText string, err error) {
	path := "/api/v2/app/setPreferences"

	content, err := json.Marshal(pref)
	if err != nil {
		return
	}
	formData := url.Values{
		"json": []string{string(content)},
	}
	err = a.api.doRequest(ctx, http.MethodPost, path, nil, formData, &respText)
	if err != nil {
		return
	}
	return
}

func (a *App) DefaultSavePath(ctx context.Context) (respText string, err error) {
	path := "/api/v2/app/defaultSavePath"
	err = a.api.doRequest(ctx, http.MethodPost, path, nil, nil, &respText)
	if err != nil {
		return
	}
	return
}
