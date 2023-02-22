package main

import (
    "log"
    "net/http"
    "encoding/json"
)

func main() {
    r, e := http.Get("http://xn--bad-tma.com:8080/foo")
    if e != nil {
        log.Fatal(e)
    }
    var v interface{}
    defer r.Body.Close()
        if e = json.NewDecoder(r.Body).Decode(&v); e != nil {
        log.Fatal(e)
    }
    log.Printf("%#v\n", v)
    return
}
