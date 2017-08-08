package main

import (
	"database/sql"
	"html/template"
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

type postResult struct {
	Result string
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/template/post.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println(err)
		tmpl.Execute(w, postResult{Result: "Failed"})
		return
	}
	content := r.PostForm.Get("oneup-content")
	log.Println(content)
	if err = insert(content); err != nil {
		log.Println(err)
		tmpl.Execute(w, postResult{Result: "Failed"})
		return
	}
	tmpl.Execute(w, postResult{Result: "Success!!"})
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
