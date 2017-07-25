package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Index is index.html templete struct
type Index struct {
	Domain string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	tpl, err := template.ParseFiles(filepath.Join(pwd, "static", "template", "index.html"))
	if err != nil {
		return
	}
	tpl.Execute(w, Index{Domain: "localhost:8080"})
}
