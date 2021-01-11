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

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

func main() {
	hotels := []hotel{
		hotel{
			"WxC",
			"Phillipine",
			"Cebu city",
			"29421",
			"Central",
		},
		hotel{
			"Aston",
			"Minnesota",
			"Elite",
			"32121",
			"Southern",
		},
		hotel{
			"Horisson",
			"7th Avenue",
			"New York",
			"12352",
			"Northern",
		},
	}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
