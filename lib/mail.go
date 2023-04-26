package lib

import (
	"encoding/base32"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"strings"

	// TODO
	// Add nicer logging
	// log "github.com/sirupsen/logrus"

	"github.com/Shopify/gomail"
)

type Mail struct {
	From         string
	To           string
	Cc           []string
	Bcc          []string
	Subject      string
	HeaderVars   []string
	EmbedFiles   []string
	AttachFiles  []string
	TemplateVars []string
	Text         string
	TextFile     string
	HtmlFile     string
	MailID       string
	CanaryURL    string
	Server       *MailServer
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// NewMail returns a new mail struct
func NewMail() *Mail {
	return &Mail{}
}

func (p *Mail) String() string {
	return fmt.Sprintf("To: %s, TemplateVars: %v, MailID: %s", p.To, p.TemplateVars, p.MailID)
}

func (p *Mail) Create() *gomail.Message {
	p.MailID = randStringRunes(8)
	m := gomail.NewMessage()

	m.SetHeader("From", p.From)
	m.SetHeader("To", p.To)

	if p.Bcc != nil {
		m.SetHeader("Bcc", p.Bcc...)
	}

	if p.Cc != nil {
		m.SetHeader("Cc", p.Cc...)
	}

	m.SetHeader("Subject", p.Subject)

	for _, hv := range p.HeaderVars {
		headers := strings.SplitN(hv, "=", 2)
		headerName := headers[0]
		headerValue := headers[1]

		m.SetHeader(headerName, headerValue)
	}

	for _, ef := range p.EmbedFiles {
		m.Embed(ef)
	}

	for _, af := range p.AttachFiles {
		m.Attach(af)
	}

	templateMap := make(map[string]string)
	for _, tv := range p.TemplateVars {

		values := strings.SplitN(tv, "=", 2)
		varName := values[0]
		varValue := values[1]

		templateMap[varName] = varValue
	}
	if p.CanaryURL != "" {

		base32MailID := base32.StdEncoding.
			WithPadding(base32.NoPadding).
			EncodeToString([]byte(p.MailID))
		templateMap["Canary"] = fmt.Sprintf("%s.G22.%s", base32MailID, p.CanaryURL)
	}

	if p.Text != "" {

		t := template.Must(template.New("message").Parse(p.Text))
		m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
			return t.Execute(w, templateMap)
		})
	}

	if p.TextFile != "" {

		t := template.Must(template.ParseFiles(p.TextFile))
		m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
			return t.Execute(w, templateMap)
		})
	}

	if p.HtmlFile != "" {
		m.SetBody("text/plain", "") // The library seems to create a text version based on the html automatically.
		t := template.Must(template.ParseFiles(p.HtmlFile))
		m.AddAlternativeWriter("text/html", func(w io.Writer) error {
			return t.Execute(w, templateMap)
		})
	}

	return m

}
