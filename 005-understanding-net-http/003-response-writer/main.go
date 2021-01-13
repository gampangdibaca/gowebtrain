package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("My-Key", "this is from me")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "whatever")
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}
