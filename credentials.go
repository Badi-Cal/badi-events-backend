package main

import (
	"log"
	"encoding/json"
)

type InstalledCredentials struct {
	ClientId string `json:"client_id"`
	ProjectId string `json:"project_id"`
	AuthUri string `json:"auth_uri"`
	TokenUri string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientSecret string `json:"client_secret"`
	RedirectUris []string `json:"redirect_uris"`
}

type Credentials struct {
	Installed InstalledCredentials `json:"installed"`
}

func CreateCredentialsJson(client_id string, client_secret string) []byte {
	var redirect_uris = []string{
		"urn:ietf:wg:oauth:2.0:oob",
		"http://localhost"}
	var installed = InstalledCredentials {
		ClientId: client_id,
		ProjectId: "test-go-1600033668852",
		AuthUri:"https://accounts.google.com/o/oauth2/auth",
		TokenUri: "https://oauth2.googleapis.com/token",
		AuthProviderX509CertUrl: "https://www.googleapis.com/oauth2/v1/certs",
		ClientSecret: client_secret,
		RedirectUris: redirect_uris}
	var credentials = Credentials {
		Installed: installed}

	b, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalf("Unable to read client secret", err)
	}
	return b
}
