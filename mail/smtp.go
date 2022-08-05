package mail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

type Letter struct {
	Address        []string // ["xxx@mail.com", "yyy@mail.com"]
	Subject        string
	Sender         string
	SenderNickName string // not required ["GoMail Robot"]
	Body           string
	BodyType       string // "text/plain" | "text/html"
}

type agent struct {
	Dialer  *gomail.Dialer
	Message *gomail.Message
}

func NewAgent(s *SMTPConfig) *agent {
	// 初始化
	agent := &agent{}

	agent.Dialer = gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	agent.Dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 解决 x509: certificate signed by unknown authority 报错问题

	return agent
}

func (d *agent) Put(l *Letter) {
	// 初始化
	d.Message = gomail.NewMessage()

	setHeader := d.Message.SetHeader
	setHeader("From", d.Message.FormatAddress(l.Sender, l.SenderNickName))
	setHeader("To", l.Address...)
	setHeader("From", l.Sender)
	setHeader("Subject", l.Subject)

	setBody := d.Message.SetBody
	setBody(l.BodyType, l.Body)
}

func (d *agent) Send() error {
	return d.Dialer.DialAndSend(d.Message)
}
