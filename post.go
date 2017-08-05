package main

import (
	"database/sql"
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

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("NO"))
		return
	}
	stmt, err := tx.Prepare("INSERT INTO t_oneup(title) VALUES(?)")
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("NO"))
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("title-1")
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("NO"))
		return
	}
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
		w.Write([]byte("NO"))
		return
	}
	w.Write([]byte("YES"))
}
