package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "github.com/jackc/pgx"
    "context"
)

type MenuStruct struct {
	Id		int
	Level		int
	Parent		int
	Fieldorder	int
	Fieldname	string
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
    var maxlevel int
    pgconfig := "user=site password=siteread host=pg.sm port=5432 dbname=spaceworld"
//    pgconfig.Host = "pg.sm"
//    pgconfig.Port = 5432
//    pgconfig.Database = "spaceworld"
//    pgconfig.User = "site"
//    pgconfig.Password = "siteread"


    conn, err := pgx.Connect(context.Background(), pgconfig)
    if err != nil {
    fmt.Println(err)
    }
    defer conn.Close(context.Background())
    err = conn.QueryRow(context.Background(), "select max(level) from catalog.mainmenu;").Scan(&maxlevel)

    inputArray := make(map[int][]MenuStruct)
        
    rows, err := conn.Query(context.Background(), "select id,level,parent,fieldorder,fieldname from catalog.mainmenu order by level")
    if err != nil {
	fmt.Println(err)
    }

    defer rows.Close()
    for rows.Next() {
	var menuItem MenuStruct
	err = rows.Scan(&menuItem.Id, &menuItem.Level, &menuItem.Parent, &menuItem.Fieldorder,&menuItem.Fieldname)
	if err != nil {
	    panic(err)
	}
	inputArray[menuItem.Level] = append(inputArray[menuItem.Level], menuItem)
    }

    fmt.Println("inputArray len = ",len(inputArray))
    
    tMenu := template.New("menu")
    tMenu, _ = tMenu.ParseFiles("templates/menu.tmpl")  // Parse template file.
    err = tMenu.Execute(w, inputArray)
    fmt.Println(err)

//    fmt.Fprintf(w, "ololo")
//    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:], menuItems)
}

func main() {
    fmt.Println("Start server")
    http.HandleFunc("/", rootHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}



