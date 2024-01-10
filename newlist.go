package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/calendar/v3"
)

type Credentials struct {
    Installed struct {
        AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
        AuthURI                 string   `json:"auth_uri"`
        ClientID                string   `json:"client_id"`
        ClientSecret            string   `json:"client_secret"`
        ProjectID               string   `json:"project_id"`
        RedirectURIs            []string `json:"redirect_uris"`
        TokenURI                string   `json:"token_uri"`
    } `json:"installed"`
}

func getClient(config *oauth2.Config) *http.Client {
    tokFile := "token.json"
    tok, err := tokenFromFile(tokFile)
    if err != nil {
        log.Fatalf("Unable to retrieve token from file: %v", err)
    }
    return config.Client(context.Background(), tok)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    tok := &oauth2.Token{}
    err = json.NewDecoder(f).Decode(tok)
    return tok, err
}

func saveToken(path string, token *oauth2.Token) {
    fmt.Printf("Saving credential file to: %s\n", path)
    f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        log.Fatalf("Unable to cache oauth token: %v", err)
    }
    defer f.Close()
    json.NewEncoder(f).Encode(token)
}

func jsonHandler(w http.ResponseWriter, r *http.Request, config *oauth2.Config) {
    tokFile := "token.json"
    if _, err := os.Stat(tokFile); os.IsNotExist(err) {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        http.Redirect(w, r, authURL, http.StatusFound)
        return
    }

    client := getClient(config)
    srv, err := calendar.New(client)
    if err != nil {
        log.Fatalf("Unable to retrieve Calendar client: %v", err)
    }

    t := time.Now().Format(time.RFC3339)
    events, err := srv.Events.List("primary").ShowDeleted(false).
        SingleEvents(true).TimeMin(t).MaxResults(30).OrderBy("startTime").Do()
    if err != nil {
        log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
    }
    json.NewEncoder(w).Encode(events)
}

func callbackHandler(config *oauth2.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")
        if code == "" {
            http.Error(w, "Missing authorization code", http.StatusBadRequest)
            return
        }
        token, err := config.Exchange(context.Background(), code)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error exchanging code for token: %v", err), http.StatusInternalServerError)
            return
        }
        saveToken("token.json", token)
        fmt.Fprintf(w, "Authorization successful! You can now close this window.")
    }
}

func main() {
    b, err := ioutil.ReadFile("credentials.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }

    config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { jsonHandler(w, r, config) })
    http.HandleFunc("/oauth2callback", callbackHandler(config))

    log.Fatal(http.ListenAndServeTLS(":8080", "/etc/letsencrypt/live/xn--bad-tma.com/fullchain.pem", "/etc/letsencrypt/live/xn--bad-tma.com/privkey.pem", nil))
}

