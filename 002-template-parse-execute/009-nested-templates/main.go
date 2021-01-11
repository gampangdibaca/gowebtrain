package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type person struct {
	FN  string
	LN  string
	Age int
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func main() {
	p1 := person{
		"James", "Bond", 42,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
