package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	"dt": date,
	"db": db,
	"mn": minus,
}

var tpl *template.Template

func firstThree(s string) string {
	return strings.TrimSpace(s)[:3]
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func date(t time.Time) string {
	return t.Format("02-01-2006")
}

func db(x float64) float64 {
	return x * 2
}

func minus(x float64) float64 {
	return x - 4
}

func main() {

	data := struct {
		FirstName string
		LastName  string
		Time      time.Time
	}{
		"James",
		"Bond",
		time.Now(),
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
