package pkg

import (
	"github.com/Xin2050/web_overwatcher/config"
	"github.com/go-gomail/gomail"
)

type Mail struct {
	Subject   string   `json:"subject"`
	PlainHtml string   `json:"plainHtml"`
	To        []string `json:"to"`
	Cc        []string `json:"cc"`
	From      string   `from:"from"`
}

func (m *Mail) Send() error {
	msg := gomail.NewMessage()
	conf := config.New().Smtp

	msg.SetHeader("To", m.To...)
	if len(m.Cc) > 0 {
		msg.SetHeader("Cc", m.Cc...)
	}
	msg.SetAddressHeader("From", conf.User, m.From)
	msg.SetBody("text/html", m.PlainHtml)
	msg.SetHeader("Subject", m.Subject)
	dialer := gomail.NewDialer(conf.Host, conf.Port, conf.User, conf.Password)
	return dialer.DialAndSend(msg)

}
