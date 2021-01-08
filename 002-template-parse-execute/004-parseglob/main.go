package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gmao"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(os.Stdout, "cat.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(os.Stdout, "poison.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
