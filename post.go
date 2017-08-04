package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Oneup struct {
	ID    int
	Title string
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./db/goneup.sqlite")
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("NO"))
		return
	}
	defer db.Close()

	sqlStmt := "SELECT id, title FROM t_oneup"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("NO"))
		return
	}
	defer rows.Close()

	oneups := []Oneup{}
	for rows.Next() {
		oneup := Oneup{}
		err = rows.Scan(&oneup.ID, &oneup.Title)
		if err != nil {
			log.Fatal(err)
		}
		oneups = append(oneups, oneup)
	}

	for _, v := range oneups {
		w.Write([]byte(fmt.Sprintf("ID:%d, Title:%s", v.ID, v.Title)))
	}
}
