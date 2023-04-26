package dummy

type Mailer struct {
}

func NewMailer() *Mailer {
	return &Mailer{}
}

func (m Mailer) Send(to string, body string, subject string) error {
	return nil
}
