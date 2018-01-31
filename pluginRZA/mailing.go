package myplugins

import (
	"fmt"
	"log"

	typeRZA "./authentication/layer2/layer3/typedef"
	gomail "gopkg.in/gomail.v2"
)

//config here
func Mail(rcp typeRZA.List, subject string, text string) {
	var list typeRZA.List
	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456") //config here
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	for _, r := range list {
		m.SetHeader("From", "no-reply@example.com")
		m.SetAddressHeader("To", r.Address, r.Name)
		m.SetHeader("Subject", "Newsletter #1")
		m.SetBody("text/html", fmt.Sprintf("Hello %s!", r.Name))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Address, err)
		}
		m.Reset()
	}
}
