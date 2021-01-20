package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
			//Secure: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
