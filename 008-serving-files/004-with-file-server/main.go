package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/holo", holo)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}

func holo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="holo.png">
	`)
}
