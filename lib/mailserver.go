package lib

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/Shopify/gomail"
)

type MailServer struct {
	SmtpServer         string
	SmtpPort           int
	SmtpUser           string
	SmtpPassword       string
	CheckTLS           bool
	MaintainConnection bool
}

func NewMailServer() *MailServer {
	return &MailServer{}
}

func (p *MailServer) SendMail(message *gomail.Message) {

	d := gomail.Dialer{Host: p.SmtpServer, Port: p.SmtpPort, Username: p.SmtpUser, Password: p.SmtpPassword}
	if p.CheckTLS == false {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if err := d.DialAndSend(message); err != nil {
		log.Print(err)
	}
}

func (p *MailServer) GetMailChannel() (chan *gomail.Message, chan bool) {

	ch := make(chan *gomail.Message)
	done := make(chan bool, 1)

	go func() {
		d := gomail.Dialer{Host: p.SmtpServer, Port: p.SmtpPort, Username: p.SmtpUser, Password: p.SmtpPassword}
		if p.CheckTLS == false {
			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		}
		// d.StartTLSPolicy = mail.MandatoryStartTLS

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(5 * time.Second):
				log.Print("No new mails received for 5s, closing channel")
				done <- true
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()

	return ch, done
}
