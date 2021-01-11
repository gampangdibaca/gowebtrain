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

type restaurant struct {
	Name string
	BM   menu
	LM   menu
	DM   menu
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
	restaurants := []restaurant{
		{"Janji", breakfast, nil, nil},
		{"Jivva", nil, lunch, dinner},
		{"Kenangan", nil, nil, menu{"Brownies"}},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
