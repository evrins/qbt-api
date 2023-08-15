package api

import (
	"context"
	"encoding/json"
	"fmt"
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

func Test_Unmarshal_nil(t *testing.T) {
	type Dog struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var s = `{"name": "evrins", "age": 42}`
	var d *Dog
	var err = json.Unmarshal([]byte(s), &d)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v\n", d)
}

func Test_Unmarshal_string(t *testing.T) {
	var s = "10240"
	var v int64
	var err = json.Unmarshal([]byte(s), &v)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
	fmt.Println(v)
}
