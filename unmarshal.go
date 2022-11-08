package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    "os"
)

type Attendees struct {
	Name        string  `json: "displayName"`
	Email       string  `json: "email"`
	ResponseStatus    string  `json: "responseStatus"`
}

type Start struct {
    DateTime    string `json: "dateTime"`
    TimeZone    string
}

type Event struct {
    ID          string
    Description string
    StartEvent       []Start `json:"start"`
    Attendee    []Attendees `json:"attendees"`
}
type Cal struct {
    Cal []Event `json:"items"`
}

func main() {
	// Open our jsonFile
    file, errf := os.Open("events.json")
    // if we os.Open returns an error then handle it
    if errf != nil {
        fmt.Println(errf)
    }
    blob, errf := ioutil.ReadAll(file)
    if errf != nil {
        fmt.Println("error:", errf)
    }

	data := Cal{}

    err := json.Unmarshal(blob, &data)
    if err != nil {
        fmt.Println("error:", err)
    }
    //fmt.Println("First array:", len(data.Cal))
    //fmt.Println(data.Cal)
	for i := 0; i < len(data.Cal); i++ {
		fmt.Println("ID: ", data.Cal[i].ID)
		fmt.Println("Description: ", data.Cal[i].Description)
		//fmt.Println("Sart Time: ", data.Cal[i].StartEvent.DateTime)
        //fmt.Println("Time Zone: ", data.Cal[i].StartTimeZone[0])
		fmt.Println("Attendees:")
        for p := 0; p < len(data.Cal[i].Attendee); p++ {
            fmt.Println("    Name: ", data.Cal[i].Attendee[p].Name)
            fmt.Println("   Email: ", data.Cal[i].Attendee[p].Email)
            //fmt.Println("Response: ", data.Cal[i].Attendee[p].ResponseStatus)
        }
	}

}
