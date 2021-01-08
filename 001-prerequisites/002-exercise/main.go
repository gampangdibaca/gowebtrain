package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println(p.fname, "Speak")
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.fname, "Kaeps")
}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}
	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(p1.lname)
	p1.pSpeak()
	fmt.Println(sa1.lname)
	sa1.saSpeak()
	sa1.pSpeak()
}
