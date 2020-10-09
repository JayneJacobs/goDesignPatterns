package main

import "fmt"

type Person struct {
	name string
}
type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// LLM close to hardware
type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildren(name string) []*Person
}

func (r *Relationships) FindAllChildren(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// HLM Business Leverl
type Research struct {
	//Break DIP
	// relationship Relationships

	// use abstraction
	browser RelationshipBrowser
}

// func (r *Research) Investigate() {
// 	relations := r.relationship.relations
// 	for _, rel := range relations {
// 		if rel.from.name == "John" && rel.relationship == Parent {
// 			fmt.Println("John has child  called ", rel.to.name)
// 		}
// 	}
// }

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildren("John") {
		fmt.Println("Johns child is ", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Mat"}

	relationships := Relationships{}
	relationships.AddParentChild(&parent, &child1)
	relationships.AddParentChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
