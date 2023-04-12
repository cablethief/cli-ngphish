/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "strings"
    "github.com/spf13/cobra"
    "gopkg.in/gomail.v2"
    "html/template"
    "io"
)

var templateVars []string


var htmlFile string
var textFile string
var text string

var subject string

var from string
var to string

var smtpServer string
var smtpPort int

var smtpUser string
var smtpPassword string

var headerVars []string
var embedFiles []string
var attachFiles []string

// spearCmd represents the spear command
var spearCmd = &cobra.Command{
    Use:   "spear",
    Short: "Target a single mailbox",
    Long: `Makes a direct connection to a SMTP server to send an email.
    
    Template may declare substitutions with {{.Name}} syntax.`,
    Run: func(cmd *cobra.Command, args []string) {
        // fmt.Println("spear called")

        m := gomail.NewMessage()

        m.SetHeader("From", from)
        m.SetHeader("To", to)

        m.SetHeader("Subject", subject)

        for _, hv := range headerVars {
            headers := strings.SplitN(hv, "=", 2)
            headerName := headers[0]
            headerValue := headers[1]
            
            m.SetHeader(headerName, headerValue)
        }

        for _, ef := range embedFiles {
            m.Embed(ef)
        }

        for _, af := range attachFiles {
            m.Attach(af)
        }

        templateMap := make(map[string]string)
        for _, tv := range templateVars {

            values := strings.SplitN(tv, "=", 2)
            varName := values[0]
            varValue := values[1]

            templateMap[varName] = varValue
        }

        if text != "" {
            
            t := template.Must(template.New("message").Parse(text))
            m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
                return t.Execute(w, templateMap)
            })
        }

        if textFile != "" {

            t := template.Must(template.ParseFiles(textFile))
            m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
                return t.Execute(w, templateMap)
            })
        }

        if htmlFile != "" {
            m.SetBody("text/plain", "") // The library seems to create a text version based on the html automatically.
            t := template.Must(template.ParseFiles(htmlFile))
            m.AddAlternativeWriter("text/html", func(w io.Writer) error {
                return t.Execute(w, templateMap)
            })
        }

        d := gomail.Dialer{Host: smtpServer, Port: smtpPort, Username: smtpUser, Password: smtpPassword}
        if err := d.DialAndSend(m); err != nil {
            panic(err)
        }


    },
}

func init() {
    rootCmd.AddCommand(spearCmd)

    spearCmd.Flags().StringVarP(&smtpServer, "server", "s", "", "SMTP server to make a connection to. eg: tenant.mail.protection.outlook.com")
    spearCmd.MarkFlagRequired("server")
    spearCmd.Flags().IntVarP(&smtpPort, "port", "p", 25, "SMTP server port to make a connection to")

    spearCmd.Flags().StringVar(&smtpUser, "username", "", "SMTP Username")
    spearCmd.Flags().StringVar(&smtpPassword, "password", "", "SMTP Password")
    spearCmd.MarkFlagsRequiredTogether("username", "password")


    spearCmd.Flags().StringVarP(&subject, "subject", "", "Good day", "Subject of the email")
    spearCmd.MarkFlagRequired("subject")


    spearCmd.Flags().StringVarP(&htmlFile, "htmltemplate", "", "", "HTML template to use")
    spearCmd.Flags().StringVarP(&textFile, "texttemplate", "", "", "TEXT template to use")
    spearCmd.Flags().StringVarP(&text, "text", "", "", "TEXT to use")

    spearCmd.MarkFlagFilename("htmltemplate","html")
    spearCmd.MarkFlagFilename("texttemplate","txt")

    spearCmd.MarkFlagsMutuallyExclusive("htmltemplate", "texttemplate", "text")

    spearCmd.Flags().StringArrayVarP(&templateVars, "templatevar", "v", []string{}, "Template variables eg: -v Name=Test -v URL=https://Test.com")

    spearCmd.Flags().StringVarP(&to, "to", "t", "", "Address to send email TO.")
    spearCmd.MarkFlagRequired("to")
    spearCmd.Flags().StringVarP(&from, "from", "f", "", "Address to send email FROM. eg: Michael <michael@testing.test>")
    spearCmd.MarkFlagRequired("from")

    spearCmd.Flags().StringArrayVarP(&headerVars, "header", "", []string{}, "Specify additional headers eg: --header Phish=Knowbe4 --header Source=example")

    spearCmd.Flags().StringArrayVarP(&embedFiles, "embed", "e", []string{}, "Specify files to embed. These can then be refrenced by their file name in the html (eg: <img src=\"cid:email-logo1.png\">) eg: --embed email-logo1.png --embed email-logo2.png")
    spearCmd.Flags().StringArrayVarP(&attachFiles, "attach", "a", []string{}, "Specify files to attach. eg: --attach test.pdf --attach average.exe")


    // TODO Error handling...
    // TODO BCC CC https://github.com/go-gomail/gomail/issues/19

    spearCmd.Flags().SortFlags = false

}
