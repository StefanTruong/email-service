package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        log.Printf("request %s /", r.Method)
        if r.Method == "POST" {
            w.WriteHeader(200)
            w.Write([]byte("Welcome to my website\n"))
            //fmt.Fprint(w, "Welcome to my website!")
        } else {
            w.WriteHeader(405)
            w.Write([]byte("method not allowed"))
        }
    })
    http.HandleFunc("/all/Hans", func (w http.ResponseWriter, r *http.Request) {
        log.Printf("request %s /", r.Method)
        if r.Method == "GET" {
            w.WriteHeader(200)
            w.Write([]byte("Deine Email ist die: hans@huhn.com\n"))
            //fmt.Fprint(w, "Welcome to my website!")
        } else {
            w.WriteHeader(405)
            w.Write([]byte("method not allowed"))
        }
    })

    http.HandleFunc("/add", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Cool! You accessed /add")
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}