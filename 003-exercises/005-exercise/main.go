package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type data struct {
	Line1 string
	Line2 string
	Line3 string
	Line4 string
	Line5 string
	Line6 string
	Line7 string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	csvFile, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var pass []data
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatalln(error)
		}
		pass = append(pass, data{
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
			line[5],
			line[6],
		},
		)
	}
	err = tpl.Execute(os.Stdout, pass)
	if err != nil {
		log.Fatalln(err)
	}
}
