package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/jackc/pgx"
    "context"
)

type MenuStruct struct {
	Id		int
	Level		int
	Parent		int
	Fieldname	string
	Fieldtype	string
	Fieldorder	int
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    var pgconfig *pgx.ConnConfig
    pgconfig.Host = "pg.sm"
    pgconfig.Port = 5432
    pgconfig.Database = "spaceworld"
    pgconfig.User = "site"
    pgconfig.Password = "siteread"


    var menuItems string

    conn, err := pgx.ConnectConfig(context.Background(), pgconfig)
    if err != nil {
    panic(err)
    }
    defer conn.Close(context.Background())

    rows, err := conn.Query(context.Background(), "select id,level,parent,fieldname,fieldtype,fieldorder from catalog.menu")
    if err != nil {
	panic(err)
    }

    defer rows.Close()

    for rows.Next() {
	var menuItem MenuStruct
	err = rows.Scan(&menuItem.Id, &menuItem.Level, &menuItem.Parent, &menuItem.Fieldname, &menuItem.Fieldtype, &menuItem.Fieldorder)
	if err != nil {
	    panic(err)
	}
	menuItems = menuItems +" | "+  menuItem.Fieldname
    }

    fmt.Fprintf(w, "ololo")
//    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:], menuItems)
}

func main() {
    fmt.Println("Start server")
    http.HandleFunc("/", rootHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}


