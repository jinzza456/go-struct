package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (Person) PrintName() { //
	fmt.Print(p.name)
}

func main() {
	var p Person

	p.name = "Smith"
	p.age = 24

	fmt.Println(p)

	p.PrintName()
}
