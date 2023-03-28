package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/vault/api"
	"golang.org/x/oauth2"
)

type Credentials struct {
	Installed struct {
		AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
		AuthURI                  string   `json:"auth_uri"`
		ClientID                 string   `json:"client_id"`
		ClientSecret             string   `json:"client_secret"`
		ProjectID                string   `json:"project_id"`
		RedirectURIs             []string `json:"redirect_uris"`
		TokenURI                 string   `json:"token_uri"`
	} `json:"installed"`
}

func main() {
	// Read the OAuth2 configuration from the credentials.json file
	file, err := os.Open("credentials.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	creds := &Credentials{}
	err = decoder.Decode(&creds)
	if err != nil {
		panic(err)
	}

	// Set up the OAuth2 configuration
	config := &oauth2.Config{
		ClientID:     creds.Installed.ClientID,
		ClientSecret: creds.Installed.ClientSecret,
		RedirectURL:  creds.Installed.RedirectURIs[0],
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  creds.Installed.AuthURI,
			TokenURL: creds.Installed.TokenURI,
		},
	}

	// Set up the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		// Exchange the authorization code for an access token
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Store the access token in Hashicorp Vault
		vault, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		secret := map[string]interface{}{
			"access_token": token.AccessToken,
		}

		path := "secret/token"

		_, err = vault.Logical().Write(path, secret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Access token stored in Hashicorp Vault: %s", token.AccessToken)
	})

	http.ListenAndServe(":8080", nil)
}

