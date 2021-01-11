package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type menu []string

func main() {
	breakfast := menu{
		"Omelette",
		"Bubur Ayam",
		"Juice",
		"Sandwich",
	}
	lunch := menu{
		"Nasi Goreng",
		"Bakso",
		"Mie Ayam",
		"Ketoprak",
		"Gado - Gado",
	}

	dinner := menu{
		"Martabak",
		"Ayam Goreng",
		"Roti Bakar",
	}
	menus := struct {
		Bf menu
		Lu menu
		Di menu
	}{
		breakfast, lunch, dinner,
	}
	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln(err)
	}
}
