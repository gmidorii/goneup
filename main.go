package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/post", postHandler)

	log.Println("run server port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
