package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/jackc/pgx"
    "github.com/jackc/pgconn"
    "context"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    var config *pgconn.Config

    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    fmt.Println("Start server")
    http.HandleFunc("/", rootHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}


