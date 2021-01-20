package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //user-id, user
var dbSessions = map[string]string{} //session-id, user-id

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: sID.String(),
			//HttpOnly: true,
			//Secure: true,
		}
		http.SetCookie(w, cookie)
	}
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		u = user{un, fn, ln}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}
	fmt.Println(cookie)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
