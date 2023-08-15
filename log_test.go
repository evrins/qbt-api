package qbt_api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestLog_MainOptions(t *testing.T) {
	opt1 := DefaultLogOptions

	opt2 := DefaultLogOptions
	opt1.Info = false
	opt2.Critical = false
	spew.Dump(opt1)
	spew.Dump(opt2)
}

func TestLog_Main(t *testing.T) {
	var opts = DefaultLogOptions
	opts.Normal = false
	logItemList, err := api.Log.Main(context.Background(), opts)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(logItemList)
}

func TestLog_Peers(t *testing.T) {
	var opts = DefaultPeerLogOptions
	peerLogItemList, err := api.Log.Peers(context.Background(), opts)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(peerLogItemList)
}
