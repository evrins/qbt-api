package api

import (
	"context"
	"log"
)

var api *Api

func init() {
	var err error
	api, err = NewApi("http://localhost:38080", EnableDebug)
	if err != nil {
		log.Fatalln(err)
	}
}

func LoginDefaultUser() {
	username := "admin"
	password := "adminadmin"
	resp, err := api.Auth.Login(context.Background(), username, password)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}
