package main

import "net/http"

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
