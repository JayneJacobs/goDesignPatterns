package main

import (
	"DesignPatterns/factories/creational"
	"DesignPatterns/factories/prototype"
	"fmt"
)

//Person interface {
type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredperson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi  my name is %s, I am %d years old\n", p.name, p.age)
}

func (p *tiredperson) SayHello() {
	fmt.Printf("Hi  my name is %s, I am %d years old\n I am too tired to talk\n", p.name, p.age)
}

// NewPerson (name string, age int) Person {
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredperson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Jayne", 129)
	p.SayHello()
	creational.Mainlike()
	prototype.Mainlike()
}

//
