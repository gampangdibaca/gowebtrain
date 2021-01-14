package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func basic(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "basic.gohtml", "this is basic path")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", "this is dog doggy dog path")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", "GAMPANGDIBACA")
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/", basic)

	http.ListenAndServe(":8080", nil)
}
