package creational

import "fmt"

// Person struct {
type Person struct {
	Name string
	Age  int
}

//factory function
// // func Newperson(name string, age int) Person {
// 	return Person{name, age}
// }

// NewPerson (name string, age int) *Person {
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

// Mainlike () {
func Mainlike() {
	//initialize directly
	p := Person{"John", 22}
	fmt.Println(p)

	// use a constructor
	p2 := NewPerson("Jayne", 29)
	p2.Age = 30
	fmt.Println(p2)
}
