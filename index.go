package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	oneups, err := selectOneup(5)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(oneups)
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	tpl, err := template.ParseFiles(filepath.Join(pwd, "static", "template", "index.html"))
	if err != nil {
		return
	}
	tpl.Execute(w, nil)
}

func selectOneup(limit int) ([]Oneup, error) {
	db, err := sql.Open("sqlite3", dbConfig)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT title FROM t_oneup ORDER BY updated_date LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	oneups := []Oneup{}
	for rows.Next() {
		oneup := Oneup{}
		err = rows.Scan(&oneup.Title)
		if err != nil {
			return nil, err
		}
		oneups = append(oneups, oneup)
	}

	return oneups, nil
}
