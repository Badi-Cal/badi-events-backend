package main

import (
<<<<<<< HEAD
    "encoding/json"
    "fmt"
    "io/ioutil"
=======
	"encoding/json"
	"fmt"
	"io/ioutil"
>>>>>>> 183f7780daac15dfb2e4cc5a1c743a5ab78680f8
    "os"
)

type Attendees struct {
<<<<<<< HEAD
    DisplayName        string  `json:"displayName"`
    Email       string  `json:"email"`
    Response    string  `json:"responseStatus"`
}

type Event struct {
    ID          string `json:"id"`
    Summary string `json:"summary"`
    Start       struct {
        DateTime    string `json:"dateTime"`
        TimeZone    string `json:"timeZone"`
    } `json:"start"`
=======
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
>>>>>>> 183f7780daac15dfb2e4cc5a1c743a5ab78680f8
    Attendee    []Attendees `json:"attendees"`
}
type Cal struct {
    Cal []Event `json:"items"`
}

func main() {
<<<<<<< HEAD
    // Open our jsonFile
=======
	// Open our jsonFile
>>>>>>> 183f7780daac15dfb2e4cc5a1c743a5ab78680f8
    file, errf := os.Open("events.json")
    // if we os.Open returns an error then handle it
    if errf != nil {
        fmt.Println(errf)
    }
    blob, errf := ioutil.ReadAll(file)
    if errf != nil {
        fmt.Println("error:", errf)
    }

<<<<<<< HEAD
    data := Cal{}
=======
	data := Cal{}
>>>>>>> 183f7780daac15dfb2e4cc5a1c743a5ab78680f8

    err := json.Unmarshal(blob, &data)
    if err != nil {
        fmt.Println("error:", err)
    }
    //fmt.Println("First array:", len(data.Cal))
    //fmt.Println(data.Cal)
<<<<<<< HEAD
    for i := 0; i < len(data.Cal); i++ {
        var item = data.Cal[i]
		fmt.Println("ID: ", item.ID)
		fmt.Println("Description: ", item.Summary)
		fmt.Println("Sart Time: ", item.Start.DateTime)
        fmt.Println("Time Zone: ", item.Start.TimeZone)
		fmt.Println("Attendees:")
        //done := false
        for p := 0; p < len(item.Attendee); p++ {
            fmt.Println("    Name: ", item.Attendee[p].DisplayName)
            fmt.Println("   Email: ", item.Attendee[p].Email)
            fmt.Println("Response: ", item.Attendee[p].Response)
            //fmt.Printf("%v", item)
            //done = true
        }
        //if done {break}
    }
=======
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
>>>>>>> 183f7780daac15dfb2e4cc5a1c743a5ab78680f8

}
