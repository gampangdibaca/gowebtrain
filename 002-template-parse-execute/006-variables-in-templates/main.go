package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

type human struct {
	FirstName string
	LastName string
}

func main() {
	h1 := human{"James","Bond"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", h1)
	if err != nil {
		log.Fatalln(err)
	}
}
