package main

import (
	"fmt"
)

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome, Apt             int
	WhatLikes                     string
	TODLikes                      string
	WithWhat                      string
}

// PersonBuilder struct
type PersonBuilder struct {
	person *Person //needs to be inited
}

// From (from string) *PersonBuilder
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (it *PersonBuilder) Build() *Person {
	return it.person
}
func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}
func (it *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*it}
}
func (it *PersonBuilder) Likes() *PersonLikesBuilder {
	return &PersonLikesBuilder{*it}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}
func (pjb *PersonJobBuilder) AsA(
	companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) Apt(
	apt int) *PersonAddressBuilder {
	it.person.Apt = apt
	return it
}
func (it *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}
func (it *PersonAddressBuilder) At(
	postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func (it *PersonAddressBuilder) WithPostcode(
	postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

type PersonLikesBuilder struct {
	PersonBuilder
}

func (it *PersonLikesBuilder) TOD(
	todlikes string) *PersonLikesBuilder {
	it.person.TODLikes = todlikes
	return it
}
func (it *PersonLikesBuilder) What(
	whatlikes string) *PersonLikesBuilder {
	it.person.WhatLikes = whatlikes
	return it
}
func (it *PersonLikesBuilder) With(
	withwhat string) *PersonLikesBuilder {
	it.person.WithWhat = withwhat
	return it
}

func main() {
	pb := NewPersonBuilder()
	person := pb.Build()
	fmt.Println(*person)
	pb.Lives().At("Grant St. ").WithPostcode("08069").In("NJ")
	pb.Likes().What("Music").With("Flute").TOD("Evening")

	fmt.Println(*person)
}
