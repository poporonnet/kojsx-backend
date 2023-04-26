package dummy

import "fmt"

type Mailer struct {
}

func NewMailer() *Mailer {
	return &Mailer{}
}

func (m Mailer) Send(to string, body string, subject string) error {
	fmt.Printf("%s\n%s\n%s\n", to, subject, body)
	return nil
}
