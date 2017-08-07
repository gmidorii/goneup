package main

import (
	"database/sql"
	"log"
	"net/http"
)

// Oneup is oneup table definition
type Oneup struct {
	Title string
}

const (
	dbConfig = "./db/goneup.sqlite"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}
	content := r.PostForm.Get("oneup-content")
	log.Println(content)
	if err = insert(content); err != nil {
		log.Println(err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/index", http.StatusOK)
}

func insert(content string) error {
	db, err := sql.Open("sqlite3", dbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO t_oneup(title) VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(content)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
