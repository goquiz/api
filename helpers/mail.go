package helpers

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	fromAddress string
	toAddresses []string
	subject     string
	body        string
	isHTML      bool
}

func NewMail(from string, to ...string) *Mail {
	return &Mail{
		fromAddress: from,
		toAddresses: to,
	}
}

func (m *Mail) Subject(s string) *Mail {
	m.subject = s
	return m
}

func (m *Mail) Body(b string, html bool) *Mail {
	m.body = b
	m.isHTML = html
	return m
}

func (m *Mail) Send() error {
	gm := gomail.NewMessage()
	gm.SetHeader("From", m.fromAddress)
	gm.SetHeader("To", m.toAddresses...)
	gm.SetHeader("Subject", m.subject)

	contentType := "text/html"
	if !m.isHTML {
		contentType = "text/plain"
	}

	gm.SetBody(contentType, m.body)

	d := gomail.NewDialer(Env.Email.SMTPHost, Env.Email.SMTPPort, Env.Email.Username, Env.Email.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(gm); err != nil {
		return err
	}

	return nil
}
