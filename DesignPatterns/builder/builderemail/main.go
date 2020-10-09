package main

import "strings"

type email struct {
	from,
	to,
	subject,
	body string
}

// EmailBuilder struct
type EmailBuilder struct {
	email email
}

// From (from string) *EmailBuilder
func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

// To populates the EmailBuilder object  with To
func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

// Subject (subject string) *EmailBuilder
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

// Body (body string) *EmailBuilder
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMaileImpl(email *email) {
	// ends the email
}

type build func(*EmailBuilder)

// SendEmail (action build)
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMaileImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("jaynejacobs@jaynejacobs").
			To("jaynejacobs@comcast.net").
			Subject("This is a subject").
			Body("This is a body")
	})
}
