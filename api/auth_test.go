package api

import (
	"context"
	"log"
	"testing"
)

func TestAuth_Login(t *testing.T) {
	resp, err := api.Auth.Login(context.Background(), "admin", "adminadmin")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}

func TestAuth_Logout(t *testing.T) {
	resp, err := api.Auth.Logout(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}
