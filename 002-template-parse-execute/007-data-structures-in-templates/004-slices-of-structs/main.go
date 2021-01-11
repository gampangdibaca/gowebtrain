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

type person struct {
	FirstName string
	LastName  string
}

func main() {
	persons := []person{
		person{"James", "Bond"},
		person{"Jason", "Bourne"},
		person{"Jack", "Williams"},
		person{"Dwayne", "Johnson"},
		person{"Brock", "Lesnar"},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", persons)
	if err != nil {
		log.Fatalln(err)
	}
}
