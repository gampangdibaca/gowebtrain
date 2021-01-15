package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func idx(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./starting-files/public")))
	http.HandleFunc("/dog/", idx)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
