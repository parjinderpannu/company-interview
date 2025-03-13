package main

import "fmt"

type person struct {
	name string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am ", p.name)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			name: "James",
		},
		ltk: true,
	}

	p1 := person{
		name: "Jenny",
	}

	// sa1.speak()
	// p1.speak()

	saySomething(sa1)
	saySomething(p1)

}
