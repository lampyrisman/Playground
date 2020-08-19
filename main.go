package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/jackc/pgx"
    "github.com/jackc/pgconn"
//    "context"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    var config *pgx.Config
    config.Host = "pg.sm"
    config.Port = 5432
    config.Database = "spaceworld"
    config.User = "site"
    config.Password = "siteread"

    conn, err := pgx.ConnectConfig(context.Background(), config)
    if err != nil {
    panic(err)
    }
    defer conn.Close()


    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    fmt.Println("Start server")
    http.HandleFunc("/", rootHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}


