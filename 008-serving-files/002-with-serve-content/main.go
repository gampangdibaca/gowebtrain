package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", holo)
	http.HandleFunc("/holo.png", holoPng)
	http.ListenAndServe(":8080", nil)
}

func holo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
<img src="holo.png">
`)
}

func holoPng(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("holo.png")
	if err != nil {
		http.Error(w, "404 error not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "404 error not found", 404)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
