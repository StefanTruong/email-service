package main

import (
    "fmt"
    "github.com/StefanTruong/email-service/db"
    "github.com/StefanTruong/email-service/env"
    "log"
    "net/http"
)

func main() {
    e, err := env.Parse()
    if err != nil {
        log.Fatal(err)
    }

    conn, err := db.Connect(e)
    defer conn.Close()

    if err != nil {
        log.Fatal(err)
    }

    // insert data into emails table (which already exists)
    _, err = conn.Exec("INSERT INTO emails (user_id, email, created_at) VALUES ('stefan', 'stefan.truong@gmail.com', current_date);")
    if err != nil {
        log.Fatal(err);
    }
    // HTTP HandlerFunction
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

    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}