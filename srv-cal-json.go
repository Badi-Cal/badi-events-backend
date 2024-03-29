package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
    "os/exec"
	"runtime"
    "time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)


	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        authCode := r.URL.Query().Get("code")
        if authCode == "" {
            http.Error(w, "Unable to read authorization code", http.StatusBadRequest)
            return
        }

        tok, err := config.Exchange(context.Background(), authCode)
        if err != nil {
            http.Error(w, "Unable to retrieve token from web", http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "Token successfully retrieved. You can close this window now.")
        return
	})

	go http.ListenAndServe(":8080", nil)

	// Open the browser to the authURL
	err := openURL(authURL)
	if err != nil {
        log.Fatalf("Unable to open browser for authURL: %v", err)
	}

	// Wait for the user to complete the authentication flow in the browser
	for {
        time.Sleep(1 * time.Second)

        resp, err := http.Get("http://localhost:8080/callback")
        if err != nil {
            continue
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            break
        }
	}

    tok, err := config.TokenSource(context.Background(), &oauth2.Token{}).Token()
	if err != nil {
		log.Fatalf("Unable to get token from TokenSource: %v", err)
	}
	return tok

}

// Helper function to open the browser with the specified URL
func openURL(url string) error {
    var err error
    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }
    return err
}

// Retrieves a token from a local file.
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

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// polls Google Calendar and writes the response to http
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
    // TimeMax(month)
    // SingleEvents(true).TimeMin(t).MaxResults(30).OrderBy("startTime").Do()
	t := time.Now().Format(time.RFC3339)
	// fnt := time.Now()
	//fortnight := time.Now().AddDate(0, 0, 14).Format(time.RFC3339)
	// month := time.Now().AddDate(0, 0, 30).Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(30).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
    json.NewEncoder(w).Encode(events)
}

// type events map[string]string

func main() {
    http.HandleFunc("/", jsonHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))

}

