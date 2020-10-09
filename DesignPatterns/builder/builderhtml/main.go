package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

// HTMLElement is...
type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat("", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat("", indentSize*(indent+1)))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

// HTMLBuilder is,,,
type HTMLBuilder struct {
	rootName string
	root     HTMLElement
}

// NewHTMLBuilder is,,,,
func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{rootName, HTMLElement{rootName, "", []HTMLElement{}}}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

// AddChild is
func (b *HTMLBuilder) AddChild(childName, childText string) {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
}

// AddChildFluent ...
func (b *HTMLBuilder) AddChildFluent(childName, childText string) *HTMLBuilder {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {

	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())
	words := []string{hello, "world"}
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("li")
		sb.WriteString(v)
		sb.WriteString("/li")
	}
	sb.WriteString("<ul>")
	fmt.Println(sb.String())

	b := NewHTMLBuilder("ul")
	b.AddChild("li", hello)
	b.AddChild("li", "world")
	fmt.Println(b.String())
	b.AddChildFluent("li", hello)
	b.AddChildFluent("li", "world")
	fmt.Println(b.String())
}
