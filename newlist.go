package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
    "runtime"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
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

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		os.Remove(tokFile) // delete the token file
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Please authenticate in the opened browser window...\n")

	// Open the URL in a browser window
    openBrowser(authURL)

	// Wait for the callback
	tokenChannel := make(chan *oauth2.Token)
	oauth2Callback := func(code string) {
		tok, err := config.Exchange(context.Background(), code)
		if err != nil {
			log.Fatalf("Unable to retrieve token from web: %v", err)
		}
		tokenChannel <- tok
	}

	http.HandleFunc("/oauth2callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		oauth2Callback(code)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	return <-tokenChannel
}

func openBrowser(url string) {
    var err error

    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        log.Fatalf("Unsupported operating system: %s", runtime.GOOS)
    }

    if err != nil {
        log.Fatalf("Failed to open browser window: %v", err)
    }
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
    fmt.Println("in the json handler.\n")
	if r.Method == "POST" {
		// Handle the POST request here
		// Example: Process the JSON request body and return a response
	} else {
		// Handle any other request methods (GET, OPTIONS, etc.) here

		// Set up the Secret Manager client
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

		t := time.Now().Format(time.RFC3339)
		events, err := srv.Events.List("primary").ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(30).OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		}
		json.NewEncoder(w).Encode(events)
	}
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
	// Read the OAuth2 configuration from the credentials.json file
	//b, err := ioutil.ReadFile("credentials.json")
	//if err != nil {
		//log.Fatalf("Unable to read client secret file: %v", err)
	//}
    // If modifying these scopes, delete your previously saved token.json.
    // config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
    //if err != nil {
        //log.Fatalf("Unable to parse client secret file to config: %v", err)
    //}


	// http.HandleFunc("/oauth2callback", callbackHandler(config))
    http.HandleFunc("/", jsonHandler)

    log.Fatal(http.ListenAndServeTLS(":8080", "/etc/letsencrypt/live/xn--bad-tma.com/fullchain.pem", "/etc/letsencrypt/live/xn--bad-tma.com/privkey.pem", nil))

}

