package qbt_api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestTransferInfo_Info(t *testing.T) {
	info, err := api.TransferInfo.Info(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(info)
}

func TestTransferInfo_SpeedLimitsMode(t *testing.T) {
	enabled, err := api.TransferInfo.SpeedLimitsMode(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(enabled)
}

func TestTransferInfo_ToggleSpeedLimitsMode(t *testing.T) {
	err := api.TransferInfo.ToggleSpeedLimitsMode(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTransferInfo_DownloadLimit(t *testing.T) {
	limit, err := api.TransferInfo.DownloadLimit(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(limit)
}

func TestTransferInfo_SetDownloadLimit(t *testing.T) {
	err := api.TransferInfo.SetDownloadLimit(context.Background(), 10240*2)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTransferInfo_UploadLimit(t *testing.T) {
	limit, err := api.TransferInfo.UploadLimit(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(limit)
}

func TestTransferInfo_SetUploadLimit(t *testing.T) {
	err := api.TransferInfo.SetUploadLimit(context.Background(), 10240*2)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTransferInfo_BanPeers(t *testing.T) {
	err := api.TransferInfo.BanPeers(context.Background(), []string{"127.0.0.1:8082"})
	if err != nil {
		log.Fatalln(err)
	}
}
