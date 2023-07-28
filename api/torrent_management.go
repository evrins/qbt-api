package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type TorrentManagement struct {
	*Api
}

type TorrentManagementInfoFilter string

const FilterAll TorrentManagementInfoFilter = "all"
const FilterDownloading TorrentManagementInfoFilter = "downloading"
const FilterSeeding TorrentManagementInfoFilter = "seeding"
const FilterCompleted TorrentManagementInfoFilter = "completed"
const FilterPaused TorrentManagementInfoFilter = "paused"
const FilterActive TorrentManagementInfoFilter = "active"
const FilterInactive TorrentManagementInfoFilter = "inactive"
const FilterResumed TorrentManagementInfoFilter = "resumed"
const FilterStalled TorrentManagementInfoFilter = "stalled"
const FilterStalledUploading TorrentManagementInfoFilter = "stalled_uploading"
const FilterStalledDownloading TorrentManagementInfoFilter = "stalled_downloading"
const FilterErrored TorrentManagementInfoFilter = "errored"

type TorrentManagementInfoOptions struct {
	Filter   TorrentManagementInfoFilter
	Category *string
	Tag      *string
	Sort     string
	Reverse  bool
	Limit    int64
	Offset   int64
	Hashes   []string
}

type TorrentManagementInfoState string

const InfoStateError TorrentManagementInfoState = "error"
const InfoStateMissingFiles TorrentManagementInfoState = "missingFiles"
const InfoStateUploading TorrentManagementInfoState = "uploading"
const InfoStatePausedUP TorrentManagementInfoState = "pausedUP"
const InfoStateQueuedUP TorrentManagementInfoState = "queuedUP"
const InfoStateStalledUP TorrentManagementInfoState = "stalledUP"
const InfoStateCheckingUP TorrentManagementInfoState = "checkingUP"
const InfoStateForcedUP TorrentManagementInfoState = "forcedUP"
const InfoStateAllocating TorrentManagementInfoState = "allocating"
const InfoStateDownloading TorrentManagementInfoState = "downloading"
const InfoStateMetaDL TorrentManagementInfoState = "metaDL"
const InfoStatePausedDL TorrentManagementInfoState = "pausedDL"
const InfoStateQueuedDL TorrentManagementInfoState = "queuedDL"
const InfoStateStalledDL TorrentManagementInfoState = "stalledDL"
const InfoStateCheckingDL TorrentManagementInfoState = "checkingDL"
const InfoStateForcedDL TorrentManagementInfoState = "forcedDL"
const InfoStateCheckingResumeData TorrentManagementInfoState = "checkingResumeData"
const InfoStateMoving TorrentManagementInfoState = "moving"
const InfoStateUnknown TorrentManagementInfoState = "unknown"

type TorrentManagementInfo struct {
	AddedOn           int                        `json:"added_on"`
	AmountLeft        int                        `json:"amount_left"`
	AutoTmm           bool                       `json:"auto_tmm"`
	Availability      int                        `json:"availability"`
	Category          string                     `json:"category"`
	Completed         int64                      `json:"completed"`
	CompletionOn      int                        `json:"completion_on"`
	ContentPath       string                     `json:"content_path"`
	DlLimit           int                        `json:"dl_limit"`
	DlSpeed           int                        `json:"dlspeed"`
	DownloadPath      string                     `json:"download_path"`
	Downloaded        int64                      `json:"downloaded"`
	DownloadedSession int                        `json:"downloaded_session"`
	Eta               int                        `json:"eta"`
	FLPiecePrio       bool                       `json:"f_l_piece_prio"`
	ForceStart        bool                       `json:"force_start"`
	Hash              string                     `json:"hash"`
	InfohashV1        string                     `json:"infohash_v1"`
	InfohashV2        string                     `json:"infohash_v2"`
	LastActivity      int                        `json:"last_activity"`
	MagnetURI         string                     `json:"magnet_uri"`
	MaxRatio          int                        `json:"max_ratio"`
	MaxSeedingTime    int                        `json:"max_seeding_time"`
	Name              string                     `json:"name"`
	NumComplete       int                        `json:"num_complete"`
	NumIncomplete     int                        `json:"num_incomplete"`
	NumLeechs         int                        `json:"num_leechs"`
	NumSeeds          int                        `json:"num_seeds"`
	Priority          int                        `json:"priority"`
	Progress          int                        `json:"progress"`
	Ratio             float64                    `json:"ratio"`
	RatioLimit        int                        `json:"ratio_limit"`
	SavePath          string                     `json:"save_path"`
	SeedingTime       int                        `json:"seeding_time"`
	SeedingTimeLimit  int                        `json:"seeding_time_limit"`
	SeenComplete      int                        `json:"seen_complete"`
	SeqDl             bool                       `json:"seq_dl"`
	Size              int64                      `json:"size"`
	State             TorrentManagementInfoState `json:"state"`
	SuperSeeding      bool                       `json:"super_seeding"`
	Tags              string                     `json:"tags"`
	TimeActive        int                        `json:"time_active"`
	TotalSize         int64                      `json:"total_size"`
	Tracker           string                     `json:"tracker"`
	TrackersCount     int                        `json:"trackers_count"`
	UpLimit           int                        `json:"up_limit"`
	Uploaded          int64                      `json:"uploaded"`
	UploadedSession   int                        `json:"uploaded_session"`
	UpSpeed           int                        `json:"upspeed"`
}

func (tm *TorrentManagement) Info(ctx context.Context, opts TorrentManagementInfoOptions) (infoList []*TorrentManagementInfo, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/info", tm.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("filter", string(opts.Filter))
	if opts.Category != nil {
		query.Add("category", *opts.Category)
	}
	if opts.Tag != nil {
		query.Add("tag", *opts.Tag)
	}
	query.Add("sort", opts.Sort)
	query.Add("reverse", strconv.FormatBool(opts.Reverse))
	query.Add("limit", strconv.FormatInt(opts.Limit, 10))
	query.Add("offset", strconv.FormatInt(opts.Offset, 10))
	query.Add("hashes", strings.Join(opts.Hashes, "|"))
	req.URL.RawQuery = query.Encode()

	resp, err := tm.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	infoList = make([]*TorrentManagementInfo, 0)
	err = json.NewDecoder(resp.Body).Decode(&infoList)
	if err != nil {
		return
	}
	return
}

type TorrentManagementProperties struct {
	AdditionDate           int     `json:"addition_date"`
	Comment                string  `json:"comment"`
	CompletionDate         int     `json:"completion_date"`
	CreatedBy              string  `json:"created_by"`
	CreationDate           int     `json:"creation_date"`
	DlLimit                int     `json:"dl_limit"`
	DlSpeed                int     `json:"dl_speed"`
	DlSpeedAvg             int     `json:"dl_speed_avg"`
	DownloadPath           string  `json:"download_path"`
	Eta                    int     `json:"eta"`
	Hash                   string  `json:"hash"`
	InfohashV1             string  `json:"infohash_v1"`
	InfohashV2             string  `json:"infohash_v2"`
	IsPrivate              bool    `json:"is_private"`
	LastSeen               int     `json:"last_seen"`
	Name                   string  `json:"name"`
	NbConnections          int     `json:"nb_connections"`
	NbConnectionsLimit     int     `json:"nb_connections_limit"`
	Peers                  int     `json:"peers"`
	PeersTotal             int     `json:"peers_total"`
	PieceSize              int     `json:"piece_size"`
	PiecesHave             int     `json:"pieces_have"`
	PiecesNum              int     `json:"pieces_num"`
	Reannounce             int     `json:"reannounce"`
	SavePath               string  `json:"save_path"`
	SeedingTime            int     `json:"seeding_time"`
	Seeds                  int     `json:"seeds"`
	SeedsTotal             int     `json:"seeds_total"`
	ShareRatio             float64 `json:"share_ratio"`
	TimeElapsed            int     `json:"time_elapsed"`
	TotalDownloaded        int64   `json:"total_downloaded"`
	TotalDownloadedSession int     `json:"total_downloaded_session"`
	TotalSize              int64   `json:"total_size"`
	TotalUploaded          int64   `json:"total_uploaded"`
	TotalUploadedSession   int     `json:"total_uploaded_session"`
	TotalWasted            int     `json:"total_wasted"`
	UpLimit                int     `json:"up_limit"`
	UpSpeed                int     `json:"up_speed"`
	UpSpeedAvg             int     `json:"up_speed_avg"`
}

func (tm *TorrentManagement) Properties(ctx context.Context, hash string) (properties *TorrentManagementProperties, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/properties", tm.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("hash", hash)
	req.URL.RawQuery = query.Encode()

	resp, err := tm.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	properties = &TorrentManagementProperties{}
	err = json.NewDecoder(resp.Body).Decode(properties)
	if err != nil {
		return
	}
	return
}

type TorrentManagementTrackerStatus int

const TrackerStatusDisabled TorrentManagementTrackerStatus = 0
const TrackerStatusNotContacted TorrentManagementTrackerStatus = 1
const TrackerStatusContractedAndWorking TorrentManagementTrackerStatus = 2
const TrackerStatusUpdating TorrentManagementTrackerStatus = 3
const TrackerStatusContractedAndNotWorking TorrentManagementTrackerStatus = 4

type TorrentManagementTracker struct {
	Msg           string `json:"msg"`
	NumDownloaded int    `json:"num_downloaded"`
	NumLeeches    int    `json:"num_leeches"`
	NumPeers      int    `json:"num_peers"`
	NumSeeds      int    `json:"num_seeds"`
	Status        int    `json:"status"`
	Tier          int    `json:"tier"`
	URL           string `json:"url"`
}

func (tm *TorrentManagement) Trackers(ctx context.Context, hash string) (trackers []*TorrentManagementTracker, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/trackers", tm.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("hash", hash)
	req.URL.RawQuery = query.Encode()

	resp, err := tm.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	trackers = make([]*TorrentManagementTracker, 0)
	err = json.NewDecoder(resp.Body).Decode(&trackers)
	if err != nil {
		return
	}
	return
}

type TorrentManagementWebSeed struct {
	Url string `json:"url"`
}

func (tm *TorrentManagement) WebSeeds(ctx context.Context, hash string) (webSeedList []*TorrentManagementWebSeed, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/webseeds", tm.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("hash", hash)
	req.URL.RawQuery = query.Encode()

	resp, err := tm.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	webSeedList = make([]*TorrentManagementWebSeed, 0)
	err = json.NewDecoder(resp.Body).Decode(&webSeedList)
	if err != nil {
		return
	}
	return
}

type TorrentManagementFilePriority int

const FilePriorityNotDownloaded TorrentManagementFilePriority = 0
const FilePriorityNormal TorrentManagementFilePriority = 1
const FilePriorityHigh TorrentManagementFilePriority = 6
const FilePriorityMax TorrentManagementFilePriority = 7

type TorrentManagementFile struct {
	Availability int                           `json:"availability"`
	Index        int                           `json:"index"`
	IsSeed       bool                          `json:"is_seed,omitempty"`
	Name         string                        `json:"name"`
	PieceRange   []int                         `json:"piece_range"`
	Priority     TorrentManagementFilePriority `json:"priority"`
	Progress     int                           `json:"progress"`
	Size         int                           `json:"size"`
}

func (tm *TorrentManagement) Files(ctx context.Context, hash string, indexes []int) (fileList []*TorrentManagementFile, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/files", tm.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}

	indexStringArray := make([]string, len(indexes))
	for i := 0; i < len(indexes); i++ {
		indexStringArray[i] = strconv.Itoa(indexes[i])
	}
	query := req.URL.Query()
	query.Add("hash", hash)
	if len(indexes) > 0 {
		query.Add("indexes", strings.Join(indexStringArray, "|"))
	}
	req.URL.RawQuery = query.Encode()

	resp, err := tm.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fileList = make([]*TorrentManagementFile, 0)
	err = json.NewDecoder(resp.Body).Decode(&fileList)
	if err != nil {
		return
	}
	return
}

type PieceState int

const PieceStateNotDownloaded PieceState = 0
const PieceStateNowDownloading PieceState = 1
const PieceStateAlreadyDownloaded PieceState = 2

func (tm *TorrentManagement) PieceStates(ctx context.Context, hash string) (pieceStates []PieceState, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/pieceStates", tm.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("hash", hash)
	req.URL.RawQuery = query.Encode()

	resp, err := tm.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	pieceStates = make([]PieceState, 0)
	err = json.NewDecoder(resp.Body).Decode(&pieceStates)
	if err != nil {
		return
	}
	return
}

func (tm *TorrentManagement) PieceHashes(ctx context.Context, hash string) (pieceHashes []string, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/pieceHashes", tm.address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("hash", hash)
	req.URL.RawQuery = query.Encode()

	resp, err := tm.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	pieceHashes = make([]string, 0)
	err = json.NewDecoder(resp.Body).Decode(&pieceHashes)
	if err != nil {
		return
	}
	return
}

func joinHashes(hashes []string, all bool) string {
	if all {
		return "all"
	}
	return strings.Join(hashes, "|")
}

func (tm *TorrentManagement) Pause(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/pause", tm.address)

	query := url.Values{}
	query.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) Resume(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/resume", tm.address)

	query := url.Values{}
	query.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) Delete(ctx context.Context, hashes []string, all bool, deleteFiles bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/delete", tm.address)

	query := url.Values{}
	query.Set("hashes", joinHashes(hashes, all))
	query.Set("deleteFiles", strconv.FormatBool(deleteFiles))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) Recheck(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/recheck", tm.address)

	query := url.Values{}
	query.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) Reannounce(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/reannounce", tm.address)

	query := url.Values{}
	query.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, query, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type TorrentManagementAddOptions struct {
	Urls               []string
	Torrents           []string
	SavePath           *string
	Cookie             *string
	Category           *string
	Tags               []string
	SkipChecking       bool
	Paused             bool
	RootFolder         bool
	Rename             *string
	UPLimit            *int64
	DLLimit            *int64
	RatioLimit         *float64
	SeedingTimeLimit   *int64
	AutoTMM            bool
	SequentialDownload bool
	FirstLastPiecePrio bool
}

func (tm *TorrentManagement) Add(ctx context.Context, opts TorrentManagementAddOptions) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/add", tm.address)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if len(opts.Urls) > 0 {
		err = writer.WriteField("urls", strings.Join(opts.Urls, "\r\n"))
		if err != nil {
			return
		}
	}

	if len(opts.Torrents) > 0 {
		for _, it := range opts.Torrents {
			var part io.Writer
			part, err = writer.CreateFormFile("torrents", filepath.Base(it))
			if err != nil {
				return
			}
			var f *os.File
			f, err = os.Open(it)
			if err != nil {
				return
			}
			_, err = io.Copy(part, f)
			if err != nil {
				return
			}
			f.Close()
		}
		err = writer.Close()
		if err != nil {
			return
		}
	}

	if opts.SavePath != nil {
		err = writer.WriteField("savepath", *opts.SavePath)
		if err != nil {
			return
		}
	}

	if opts.Cookie != nil {
		err = writer.WriteField("cookie", *opts.Cookie)
		if err != nil {
			return
		}
	}

	if opts.Category != nil {
		err = writer.WriteField("category", *opts.Category)
		if err != nil {
			return
		}
	}

	if len(opts.Tags) != 0 {
		err = writer.WriteField("tags", strings.Join(opts.Tags, "|"))
		if err != nil {
			return
		}
	}

	err = writer.WriteField("skip_checking", strconv.FormatBool(opts.SkipChecking))
	if err != nil {
		return
	}

	err = writer.WriteField("paused", strconv.FormatBool(opts.Paused))
	if err != nil {
		return
	}

	err = writer.WriteField("root_folder", strconv.FormatBool(opts.RootFolder))
	if err != nil {
		return
	}

	if opts.Rename != nil {
		err = writer.WriteField("rename", *opts.Rename)
		if err != nil {
			return
		}
	}

	if opts.UPLimit != nil {
		err = writer.WriteField("upLimit", strconv.FormatInt(*opts.UPLimit, 10))
		if err != nil {
			return
		}
	}

	if opts.DLLimit != nil {
		err = writer.WriteField("dlLimit", strconv.FormatInt(*opts.DLLimit, 10))
		if err != nil {
			return
		}
	}

	if opts.RatioLimit != nil {
		err = writer.WriteField("ratioLimit", strconv.FormatFloat(*opts.RatioLimit, 'f', 0, 64))
		if err != nil {
			return
		}
	}

	if opts.SeedingTimeLimit != nil {
		err = writer.WriteField("seedingTimeLimit", strconv.FormatInt(*opts.SeedingTimeLimit, 10))
		if err != nil {
			return
		}
	}

	err = writer.WriteField("autoTMM", strconv.FormatBool(opts.AutoTMM))
	if err != nil {
		return
	}

	err = writer.WriteField("sequentialDownload", strconv.FormatBool(opts.SequentialDownload))
	if err != nil {
		return
	}
	err = writer.WriteField("firstLastPiecePrio", strconv.FormatBool(opts.FirstLastPiecePrio))
	if err != nil {
		return
	}

	resp, _, err := tm.doRequestWithMultiPartForm(ctx, http.MethodPost, link, writer.FormDataContentType(), nil, body)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) AddTrackers(ctx context.Context, hash string, urls []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/addTrackers", tm.address)

	formData := url.Values{}
	formData.Set("hash", hash)
	formData.Set("urls", strings.Join(urls, "\n"))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) EditTracker(ctx context.Context, hash, originUrl, newUrl string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/editTracker", tm.address)

	formData := url.Values{}
	formData.Set("hash", hash)
	formData.Set("origUrl", originUrl)
	formData.Set("newUrl", newUrl)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) RemoveTrackers(ctx context.Context, hash string, urls []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/removeTrackers", tm.address)

	formData := url.Values{}
	formData.Set("hash", hash)
	formData.Set("urls", strings.Join(urls, "|"))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type AddPeerResult struct {
	Added  int64 `json:"added"`
	Failed int64 `json:"failed"`
}

type AddPeerResponse map[string]AddPeerResult

func (tm *TorrentManagement) AddPeers(ctx context.Context, hashes, peers []string) (addPeerResponse AddPeerResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/addPeers", tm.address)

	formData := url.Values{}
	formData.Set("hashes", strings.Join(hashes, "|"))
	formData.Set("peers", strings.Join(peers, "|"))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	addPeerResponse = map[string]AddPeerResult{}
	err = json.NewDecoder(resp).Decode(&addPeerResponse)
	if err != nil {
		return
	}
	return
}

func (tm *TorrentManagement) IncreasePriority(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/increasePrio", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) DecreasePriority(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/decreasePrio", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) TopPriority(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/topPrio", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) BottomPriority(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/bottomPrio", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) SetFilePriority(ctx context.Context, hash string, ids []int, priority TorrentManagementFilePriority) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/filePrio", tm.address)

	stringIdArray := make([]string, len(ids))
	for i, it := range ids {
		stringIdArray[i] = strconv.Itoa(it)
	}
	formData := url.Values{}
	formData.Set("hash", hash)
	formData.Set("id", strings.Join(stringIdArray, "|"))
	formData.Set("priority", strconv.Itoa(int(priority)))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type DownloadLimitResponse map[string]int

func (tm *TorrentManagement) DownloadLimit(ctx context.Context, hashes []string, all bool) (downloadLimitResponse DownloadLimitResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/downloadLimit", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	downloadLimitResponse = map[string]int{}
	err = json.NewDecoder(resp).Decode(&downloadLimitResponse)
	if err != nil {
		return
	}
	return
}

func (tm *TorrentManagement) SetDownloadLimit(ctx context.Context, hashes []string, all bool, limit int) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setDownloadLimit", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("limit", strconv.Itoa(limit))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

// SetShareLimits
// @ratioLimit -2 use global limit, -1 no limit
// @seedingTimeLimit -2 use global limit, -1 no limit
func (tm *TorrentManagement) SetShareLimits(ctx context.Context, hashes []string, all bool, ratioLimit float64, seedingTimeLimit int64) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setShareLimits", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("ratioLimit", strconv.FormatFloat(ratioLimit, 'f', 0, 64))
	formData.Set("seedingTimeLimit", strconv.FormatInt(seedingTimeLimit, 10))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type UploadLimitResponse map[string]int

func (tm *TorrentManagement) UploadLimit(ctx context.Context, hashes []string, all bool) (uploadLimitResponse UploadLimitResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/uploadLimit", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	uploadLimitResponse = map[string]int{}
	err = json.NewDecoder(resp).Decode(&uploadLimitResponse)
	if err != nil {
		return
	}
	return
}

func (tm *TorrentManagement) SetUploadLimit(ctx context.Context, hashes []string, all bool, limit int64) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setUploadLimit", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("limit", strconv.FormatInt(limit, 10))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) SetLocation(ctx context.Context, hashes []string, all bool, location string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setLocation", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("location", location)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) Rename(ctx context.Context, hash, name string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/rename", tm.address)

	formData := url.Values{}
	formData.Set("hash", hash)
	formData.Set("name", name)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) SetCategory(ctx context.Context, hashes []string, all bool, category string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setCategory", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("category", category)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

type CategoryResponse map[string]Category

func (tm *TorrentManagement) Categories(ctx context.Context) (categoryResponse CategoryResponse, err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/categories", tm.address)

	resp, _, err := tm.doRequest(ctx, http.MethodGet, link, nil, nil)
	if err != nil {
		return
	}
	defer resp.Close()
	categoryResponse = map[string]Category{}
	err = json.NewDecoder(resp).Decode(&categoryResponse)
	return
}

func (tm *TorrentManagement) CreateCategory(ctx context.Context, category *Category) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/createCategory", tm.address)

	formData := url.Values{}
	formData.Set("category", category.Name)
	formData.Set("savePath", category.SavePath)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) EditCategory(ctx context.Context, category *Category) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/editCategory", tm.address)

	formData := url.Values{}
	formData.Set("category", category.Name)
	formData.Set("savePath", category.SavePath)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) RemoveCategories(ctx context.Context, categories []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/removeCategories", tm.address)

	formData := url.Values{}
	formData.Set("categories", strings.Join(categories, "\n"))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) AddTags(ctx context.Context, hashes []string, all bool, tags []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/addTags", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("tags", strings.Join(tags, ","))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) RemoveTags(ctx context.Context, hashes []string, all bool, tags []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/removeTags", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("tags", strings.Join(tags, ","))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) CreateTags(ctx context.Context, tags []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/createTags", tm.address)

	formData := url.Values{}
	formData.Set("tags", strings.Join(tags, ","))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) DeleteTags(ctx context.Context, tags []string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/deleteTags", tm.address)

	formData := url.Values{}
	formData.Set("tags", strings.Join(tags, ","))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) SetAutoManagement(ctx context.Context, hashes []string, all bool, enable bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setAutoManagement", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("enable", strconv.FormatBool(enable))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) ToggleSequentialDownload(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/toggleSequentialDownload", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) ToggleFirstLastPiecePriority(ctx context.Context, hashes []string, all bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/toggleFirstLastPiecePrio", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) SetForceStart(ctx context.Context, hashes []string, all bool, forceStart bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setForceStart", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("value", strconv.FormatBool(forceStart))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) SetSuperSeeding(ctx context.Context, hashes []string, all bool, superSeeding bool) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/setSuperSeeding", tm.address)

	formData := url.Values{}
	formData.Set("hashes", joinHashes(hashes, all))
	formData.Set("value", strconv.FormatBool(superSeeding))

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) RenameFile(ctx context.Context, hash, oldPath, newPath string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/renameFile", tm.address)

	formData := url.Values{}
	formData.Set("hash", hash)
	formData.Set("oldPath", oldPath)
	formData.Set("newPath", newPath)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}

func (tm *TorrentManagement) RenameFolder(ctx context.Context, hash, oldPath, newPath string) (err error) {
	link := fmt.Sprintf("%s/api/v2/torrents/renameFolder", tm.address)

	formData := url.Values{}
	formData.Set("hash", hash)
	formData.Set("oldPath", oldPath)
	formData.Set("newPath", newPath)

	resp, _, err := tm.doRequest(ctx, http.MethodPost, link, nil, formData)
	if err != nil {
		return
	}
	defer resp.Close()
	return
}
