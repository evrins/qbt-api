package api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestSync_MainData(t *testing.T) {
	resp, err := api.Sync.MainData(context.Background(), 0)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}

func TestSync_TorrentPeers(t *testing.T) {
	resp, err := api.Sync.TorrentPeers(context.Background(), "c697e22d8b385a4a667d773467a840adae200919", 0)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(resp)
}
