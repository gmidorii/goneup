package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Oneup is oneup table definition
type Oneup struct {
	Title       string
	CreatedDate string
	UpdatedDate string
}

type postResult struct {
	Result string
	Title  string
	Date   string
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
	title := r.PostForm.Get("oneup-content")
	date := time.Now()
	log.Println(title)
	oneup := Oneup{
		Title:       title,
		CreatedDate: date.Format(dateLayout),
		UpdatedDate: date.Format(dateLayout),
	}
	if err = insert(oneup); err != nil {
		log.Println(err)
		tmpl.Execute(w, postResult{Result: "Failed"})
		return
	}
	tmpl.Execute(w, postResult{Result: "Success!!", Title: title, Date: date.Format(dateLayout)})
}

func insert(oneup Oneup) error {
	db, err := sql.Open("sqlite3", dbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO t_oneup(title, created_date, updated_date) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(oneup.Title, oneup.CreatedDate, oneup.UpdatedDate)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
