package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	tpl, err := template.ParseFiles(filepath.Join(pwd, "static", "template", "index.html"))
	if err != nil {
		return
	}
	tpl.Execute(w, nil)
}
