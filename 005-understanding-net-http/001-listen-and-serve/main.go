package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "anything that can fit here")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
