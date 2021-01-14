package main

import (
	"io"
	"net/http"
)

func basic(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is basic path")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is dog doggy dog path")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GAMPANGDIBACA")
}

func main() {
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/", basic)

	http.ListenAndServe(":8080", nil)
}
