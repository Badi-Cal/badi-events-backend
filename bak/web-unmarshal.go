package main

import (
    "encoding/json"
    "fmt"
	"net/http"
    "io/ioutil"
    "os"
)

type Attendees struct {
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
    Attendee    []Attendees `json:"attendees"`
}
type Cal struct {
    Cal []Event `json:"items"`
}

func main() {
	r, e := http.Get("http://xn--bad-tma.com:8088/foo")
    if e != nil {
        log.Fatal(e)
    }
    //var v interface{}
    //defer r.Body.Close()
        //if e = json.NewDecoder(r.Body).Decode(&v); e != nil {
        //log.Fatal(e)
    //}
    //log.Printf("%#v\n", v)

    data := Cal{}

    err := json.Unmarshal(r, &data)
    if err != nil {
        fmt.Println("error:", err)
    }
    //fmt.Println("First array:", len(data.Cal))
    //fmt.Println(data.Cal)
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
    return

}
