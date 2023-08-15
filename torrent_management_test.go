package qbt_api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestTorrentManagement_Info(t *testing.T) {
	var opts = TorrentManagementInfoOptions{
		Filter:   FilterAll,
		Category: nil,
		Tag:      nil,
		Sort:     "",
		Reverse:  false,
		Limit:    0,
		Offset:   0,
		Hashes:   nil,
	}
	infoList, err := api.TorrentManagement.Info(context.Background(), opts)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(infoList)
}

func TestTorrentManagement_Properties(t *testing.T) {
	hash := "c697e22d8b385a4a667d773467a840adae200919"
	resp, err := api.TorrentManagement.Properties(context.Background(), hash)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_Trackers(t *testing.T) {
	hash := "c697e22d8b385a4a667d773467a840adae200919"
	resp, err := api.TorrentManagement.Trackers(context.Background(), hash)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_WebSeeds(t *testing.T) {
	hash := "c697e22d8b385a4a667d773467a840adae200919"
	resp, err := api.TorrentManagement.WebSeeds(context.Background(), hash)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_Files(t *testing.T) {
	hash := "c697e22d8b385a4a667d773467a840adae200919"
	resp, err := api.TorrentManagement.Files(context.Background(), hash, []int{})
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_PieceStates(t *testing.T) {
	hash := "c697e22d8b385a4a667d773467a840adae200919"
	resp, err := api.TorrentManagement.PieceStates(context.Background(), hash)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_PieceHashes(t *testing.T) {
	hash := "c697e22d8b385a4a667d773467a840adae200919"
	resp, err := api.TorrentManagement.PieceHashes(context.Background(), hash)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_Pause(t *testing.T) {
	hashes := []string{"c697e22d8b385a4a667d773467a840adae200919"}
	err := api.TorrentManagement.Pause(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_Add_URLs(t *testing.T) {
	sp := "/downloads"
	opts := TorrentManagementAddOptions{
		Urls: []string{
			"https://www.btbtt13.com/attach-download-fid-953-aid-6054931.htm",
		},
		SavePath: &sp,
	}
	var err = api.TorrentManagement.Add(context.Background(), opts)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_Add_Files(t *testing.T) {
	sp := "/downloads"
	opts := TorrentManagementAddOptions{
		Torrents: []string{
			"./torrent.torrent",
			"./torrent2.torrent",
		},
		SavePath: &sp,
	}
	err := api.TorrentManagement.Add(context.Background(), opts)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_AddTrackers(t *testing.T) {
	var hash = "c697e22d8b385a4a667d773467a840adae200919"
	var urls = []string{
		"udp://thouvenin.cloud:6969/announce",
		"ws://hub.bugout.link:80/announce",
	}
	var err = api.TorrentManagement.AddTrackers(context.Background(), hash, urls)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_EditTracker(t *testing.T) {
	var hash = "c697e22d8b385a4a667d773467a840adae200919"

	var originUrl = "udp://thouvenin.cloud:6969/announce"
	var newUrl = "udp://thouvenin.cloud:16969/announce"
	var err = api.TorrentManagement.EditTracker(context.Background(), hash, originUrl, newUrl)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_RemoveTrackers(t *testing.T) {
	var hash = ""
	var urls = []string{
		"udp://thouvenin.cloud:16969/announce",
		"ws://hub.bugout.link:80/announce",
	}
	var err = api.TorrentManagement.RemoveTrackers(context.Background(), hash, urls)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_AddPeers(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var urls = []string{
		"127.0.0.1:9090",
	}

	var resp, err = api.TorrentManagement.AddPeers(context.Background(), hashes, urls)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_IncreasePriority(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var err = api.TorrentManagement.IncreasePriority(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_DecreasePriority(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var err = api.TorrentManagement.DecreasePriority(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_TopPriority(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var err = api.TorrentManagement.TopPriority(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_BottomPriority(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var err = api.TorrentManagement.BottomPriority(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_SetFilePriority(t *testing.T) {
	var hash = "c697e22d8b385a4a667d773467a840adae200919"
	var ids = []int{0, 1}
	var err = api.TorrentManagement.SetFilePriority(context.Background(), hash, ids, FilePriorityMax)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_DownloadLimit(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var resp, err = api.TorrentManagement.DownloadLimit(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_SetDownloadLimit(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var limit = 10240
	var err = api.TorrentManagement.SetDownloadLimit(context.Background(), hashes, false, limit)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_SetShareLimits(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var err = api.TorrentManagement.SetShareLimits(context.Background(), hashes, false, 1.0, 3600)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_UploadLimit(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var resp, err = api.TorrentManagement.UploadLimit(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_SetUploadLimit(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var err = api.TorrentManagement.SetUploadLimit(context.Background(), hashes, false, 10240)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_Rename(t *testing.T) {
	var hash = "c697e22d8b385a4a667d773467a840adae200919"
	var name = "some new name"

	var err = api.TorrentManagement.Rename(context.Background(), hash, name)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_SetCategory(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var category = "c2"

	var err = api.TorrentManagement.SetCategory(context.Background(), hashes, false, category)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_Categories(t *testing.T) {
	var resp, err = api.TorrentManagement.Categories(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestTorrentManagement_CreateCategory(t *testing.T) {
	c := &Category{
		Name:     "c3",
		SavePath: "/c3",
	}

	var err = api.TorrentManagement.CreateCategory(context.Background(), c)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_EditCategory(t *testing.T) {
	c := &Category{
		Name:     "c3",
		SavePath: "/c33",
	}

	var err = api.TorrentManagement.EditCategory(context.Background(), c)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_RemoveCategories(t *testing.T) {
	var categories = []string{"c3", "c2"}

	var err = api.TorrentManagement.RemoveCategories(context.Background(), categories)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_AddTags(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var tags = []string{
		"11", "22",
	}

	var err = api.TorrentManagement.AddTags(context.Background(), hashes, false, tags)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_RemoveTags(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}
	var tags = []string{
		"11", "22",
	}

	var err = api.TorrentManagement.RemoveTags(context.Background(), hashes, false, tags)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_CreateTags(t *testing.T) {
	var tags = []string{
		"111", "222",
	}

	var err = api.TorrentManagement.CreateTags(context.Background(), tags)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_DeleteTags(t *testing.T) {
	var tags = []string{
		"111", "222",
	}

	var err = api.TorrentManagement.DeleteTags(context.Background(), tags)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_SetAutoManagement(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var err = api.TorrentManagement.SetAutoManagement(context.Background(), hashes, false, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_ToggleSequentialDownload(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var err = api.TorrentManagement.ToggleSequentialDownload(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_ToggleFirstLastPiecePriority(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var err = api.TorrentManagement.ToggleFirstLastPiecePriority(context.Background(), hashes, false)
	if err != nil {
		log.Fatalln(err)
	}
}
func TestTorrentManagement_SetForceStart(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var err = api.TorrentManagement.SetForceStart(context.Background(), hashes, false, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_SetSuperSeeding(t *testing.T) {
	var hashes = []string{
		"c697e22d8b385a4a667d773467a840adae200919",
	}

	var err = api.TorrentManagement.SetSuperSeeding(context.Background(), hashes, false, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_RenameFile(t *testing.T) {
	var hash = "c697e22d8b385a4a667d773467a840adae200919"

	var oldPath = "【高清剧集网发布 www.DDHDTV.com】魔术士欧菲流浪之旅 基姆拉克篇[全11集][中文字幕].Majutsushi.Orphen.Hagure.Tabi.Kimluck.Hen.2021.S02.Complete.1080p.NF.WEB-DL.x264.DDP2.0-Huawei/Majutsushi.Orphen.Hagure.Tabi.Kimluck.Hen.2021.S02E02.1080p.NF.WEB-DL.x264.DDP2.0-Huawei.mkv"
	var newPath = "【高清剧集网发布 www.DDHDTV.com】魔术士欧菲流浪之旅 基姆拉克篇[全11集][中文字幕].Majutsushi.Orphen.Hagure.Tabi.Kimluck.Hen.2021.S02.Complete.1080p.NF.WEB-DL.x264.DDP2.0-Huawei/s02e02.mkv"

	var err = api.TorrentManagement.RenameFile(context.Background(), hash, oldPath, newPath)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTorrentManagement_RenameFolder(t *testing.T) {
	var hash = "c697e22d8b385a4a667d773467a840adae200919"

	var oldPath = "【高清剧集网发布 www.DDHDTV.com】魔术士欧菲流浪之旅 基姆拉克篇[全11集][中文字幕].Majutsushi.Orphen.Hagure.Tabi.Kimluck.Hen.2021.S02.Complete.1080p.NF.WEB-DL.x264.DDP2.0-Huawei"
	var newPath = "s02"

	var err = api.TorrentManagement.RenameFolder(context.Background(), hash, oldPath, newPath)
	if err != nil {
		log.Fatalln(err)
	}
}
