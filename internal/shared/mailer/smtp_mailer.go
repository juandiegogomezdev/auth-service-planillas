package mailer

type SMTPMailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewSMTPMailer(host string, port int, username, password, from string) *SMTPMailer {
	return &SMTPMailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (m *SMTPMailer) Send(to, subject, body string) error {
	return nil
}
