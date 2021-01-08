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

type human interface {
	Speak()
}

func (p person) Speak() {
	fmt.Println(p.fname, "Speak")
}

func (sa secretAgent) Speak() {
	fmt.Println(sa.fname, "Kaeps")
}

func info(h human) {
	h.Speak()
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
	p1.Speak()
	fmt.Println(sa1.lname)
	sa1.Speak()
	sa1.person.Speak()
	fmt.Println("-----separator-----")
	info(p1)
	info(sa1)
}
