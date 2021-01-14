package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog hotdog dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat hotcat cat")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
