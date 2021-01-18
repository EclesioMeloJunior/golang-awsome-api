package notification

import (
	"errors"
	"fmt"
	"go-challenge/config"
	"net/smtp"
)

// Emailer interface abstracts the implementation
// and allow the dependency inversion
type Emailer interface {
	Send() error
	Informations(to, subject string, body []byte)
}

type email struct {
	from     string
	password string
	server   string
	port     int

	to      string
	subject string
	body    []byte
}

// NewEmailer returns an implementation of Emailer
// using the golang lib net/smtp
func NewEmailer(c *config.Config) Emailer {
	return &email{
		from:     c.EmailUser,
		port:     c.EmailPort,
		password: c.EmailPass,
		server:   c.EmailServer,
	}
}

func (e *email) Informations(to, subject string, body []byte) {
	e.to = to
	e.body = body
	e.subject = subject
}

func (e *email) Send() error {
	if len(e.body) < 1 || e.to == "" || e.subject == "" {
		return errors.New("You need to fill the email send informations")
	}

	auth := smtp.PlainAuth(
		"", e.from, e.password, e.server,
	)

	addr := fmt.Sprintf("%s:%v", e.server, e.port)

	return smtp.SendMail(
		addr, auth, e.from, []string{e.to}, e.body,
	)
}
