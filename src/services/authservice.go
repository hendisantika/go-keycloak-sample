package services

import (
	"github.com/Nerzal/gocloak/v13"
	"os"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

var (
	clientId     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	realm        = os.Getenv("REALM")
	hostname     = os.Getenv("HOST")
)

var client gocloak.GoCloak

func InitializeOauthServer() {
	client = gocloak.NewClient(hostname)
}
