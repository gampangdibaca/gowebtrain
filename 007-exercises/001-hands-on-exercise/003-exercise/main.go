package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func basic(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "basic.gohtml", "this is basic path")
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", "this is dog doggy dog path")
	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "me.gohtml", "GAMPANGDIBACA")
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/me/", http.HandlerFunc(me))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/", http.HandlerFunc(basic))

	http.ListenAndServe(":8080", nil)
}
