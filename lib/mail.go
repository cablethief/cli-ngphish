package lib

import (
	"html/template"
	"io"
	"strings"

	"gopkg.in/gomail.v2"
)

type MailStruct struct {
	from         string
	to           string
	subject      string
	headerVars   []string
	embedFiles   []string
	attachFiles  []string
	templateVars []string
	text         string
	textFile     string
	htmlFile     string
	server       *ServerDetails
}

type ServerDetails struct {
	smtpServer   string
	smtpPort     int
	smtpUser     string
	smtpPassword string
}

func (p *MailStruct) sendMail() {

	m := gomail.NewMessage()

	m.SetHeader("From", p.from)
	m.SetHeader("To", p.to)

	m.SetHeader("Subject", p.subject)

	for _, hv := range p.headerVars {
		headers := strings.SplitN(hv, "=", 2)
		headerName := headers[0]
		headerValue := headers[1]

		m.SetHeader(headerName, headerValue)
	}

	for _, ef := range p.embedFiles {
		m.Embed(ef)
	}

	for _, af := range p.attachFiles {
		m.Attach(af)
	}

	templateMap := make(map[string]string)
	for _, tv := range p.templateVars {

		values := strings.SplitN(tv, "=", 2)
		varName := values[0]
		varValue := values[1]

		templateMap[varName] = varValue
	}

	if p.text != "" {

		t := template.Must(template.New("message").Parse(p.text))
		m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
			return t.Execute(w, templateMap)
		})
	}

	if p.textFile != "" {

		t := template.Must(template.ParseFiles(p.textFile))
		m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
			return t.Execute(w, templateMap)
		})
	}

	if p.htmlFile != "" {
		m.SetBody("text/plain", "") // The library seems to create a text version based on the html automatically.
		t := template.Must(template.ParseFiles(p.htmlFile))
		m.AddAlternativeWriter("text/html", func(w io.Writer) error {
			return t.Execute(w, templateMap)
		})
	}

	d := gomail.Dialer{Host: p.server.smtpServer, Port: p.server.smtpPort, Username: p.server.smtpUser, Password: p.server.smtpPassword}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
