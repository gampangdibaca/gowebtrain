package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ParseFiles("poison.gmao", "cat.gmao")
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
