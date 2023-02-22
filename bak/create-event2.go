package main

import (
        "context"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"

        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/calendar/v3"
        "google.golang.org/api/option"
)

type Attendee struct {
    DisplayName string
    Email       string
}
type End struct {
    DateTime    string
    TimeZone    string
}
type Start struct {
    DateTime    string
    TimeZone    string
}
type NewEvent struct {
    Attendees   []Attendee
    Description string
    Ends        End
    Kind        string
    Location    string
    Starts      Start
    Summary     string
}

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

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
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

// create event API
func handlePost(w http.ResponseWriter, r *http.Request) {
	var event NewEvent
    ctx := context.Background()
    b, err := ioutil.ReadFile("create-event-credentials.json")
    if err != nil {
            log.Fatalf("Unable to read client secret file: %v", err)
    }
     // scope, is broad See, edit, share, and permanently delete all the calendars
    config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
    if err != nil {
            log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(config)

    err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    event:= calendar.Event{}
    srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
            log.Fatalf("Unable to retrieve Calendar client: %v", err)
    }

    calendarId := "primary"
    serializedEvent, err = srv.Events.Insert(calendarId, serializedEevent).Do()
    if err != nil {
        log.Fatalf("Unable to create event. %v\n", err)
    }
    fmt.Printf("Event created: %s\n", event.HtmlLink)
	fmt.Println(event.Summary)
	fmt.Println(event.Description)
	fmt.Println(event.Kind)
	fmt.Println(event.Location)
	fmt.Println(event.Starts)
	fmt.Println(event.Ends)
	fmt.Println(event.Attendees)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received person data"))
}

func main() {
	http.HandleFunc("/create", handlePost)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
    fmt.Println(err)
    }


}

