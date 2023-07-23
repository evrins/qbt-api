package api

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func init() {
	LoginDefaultUser()
}

func TestApp_Version(t *testing.T) {
	resp, err := api.App.Version(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}

func TestApp_WebApiVersion(t *testing.T) {
	resp, err := api.App.WebApiVersion(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}

func TestApp_BuildInfo(t *testing.T) {
	bi, err := api.App.BuildInfo(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(bi)
}

func TestApp_Shutdown(t *testing.T) {
	respText, err := api.App.Shutdown(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(respText)
}

func TestApp_Preferences(t *testing.T) {
	pref, err := api.App.Preferences(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(pref)
}

func TestApp_DefaultSavePath(t *testing.T) {
	respText, err := api.App.DefaultSavePath(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(respText)
}
