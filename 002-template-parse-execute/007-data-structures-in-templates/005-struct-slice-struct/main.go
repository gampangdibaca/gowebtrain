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

type group struct {
	GroupMember []person
	GroupName   string
}

func main() {
	persons := []person{
		person{"James", "Bond"},
		person{"Jason", "Bourne"},
		person{"Jack", "Williams"},
		person{"Dwayne", "Johnson"},
		person{"Brock", "Lesnar"},
	}
	gr := group{persons, "Cool Guys"}
	// more efficient way
	//gr := struct {
	//	GroupMember []person
	//	GroupName   string
	//}{persons, "Cool Guys"}
	// and we can delete group type at the top
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", gr)
	if err != nil {
		log.Fatalln(err)
	}
}
