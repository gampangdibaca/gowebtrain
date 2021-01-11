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

func main() {

	//names := []string{"Budi", "Charlie", "David", "Evan", "Fauzan"}
	names := map[string]string{
		"Nimo":     "omiN",
		"Youtube":  "ebutuoY",
		"Twitch":   "hctiwT",
		"Mixer":    "rexiM",
		"Facebook": "koobecaF",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}
}
