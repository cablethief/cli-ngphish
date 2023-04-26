/*
Copyright © 2022 Michael Kruger @_Cablethief
*/
package cmd

import (
	"log"

	"github.com/cablethief/cli-ngphish/lib"

	"github.com/spf13/cobra"
)

var csvFile string
var mailspray = &lib.Mail{}

// Always needs a To and a From
//
// sprayCmd represents the spray command
var sprayCmd = &cobra.Command{
	Use:   "spray",
	Short: "Spray at multiple mailboxes based off a CSV",
	Long: `Sprays multiple mailboxes pulling the To, From and template variables from a CSV. For example:

The CSV will always require a "To" header, the rest of the variables in a 
template may be specified with other headers. The normal GoPhish group CSV would
be "To, FirstName, LastName, Position" where you can specify {{.FirstName}} etc in the template.`,
	Run: func(cmd *cobra.Command, args []string) {
		EmailsParsed := lib.ParsePhishCSV(csvFile)

		if !MailServer.MaintainConnection {
			log.Println("Starting Sender Routine")
			mailpipe, done := MailServer.GetMailChannel()
			for i, _ := range EmailsParsed {
				mailspray.To = EmailsParsed[i].To
				mailspray.TemplateVars = EmailsParsed[i].TemplateVars

				mailpipe <- mailspray.Create()
				log.Println(mailspray)

			}
			<-done
			log.Println("Closing Sender Routine")
			close(mailpipe)
			close(done)

		} else {
			for i, _ := range EmailsParsed {
				mailspray.To = EmailsParsed[i].To
				mailspray.TemplateVars = EmailsParsed[i].TemplateVars

				log.Println(mailspray)
				MailServer.SendMail(mailspray.Create())
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(sprayCmd)

	// Make sure to do a check to ensure the CSV doesn't include a from
	sprayCmd.Flags().StringVarP(&mailspray.From, "from", "f", "", "Address to send email FROM. eg: Michael <michael@testing.test>")

	sprayCmd.Flags().StringVarP(&mailspray.Subject, "subject", "", "Good day", "Subject of the email")
	sprayCmd.MarkFlagRequired("subject")

	sprayCmd.Flags().StringVarP(&mailspray.HtmlFile, "htmltemplate", "", "", "HTML template to use")
	sprayCmd.Flags().StringVarP(&mailspray.TextFile, "texttemplate", "", "", "TEXT template to use")
	sprayCmd.Flags().StringVarP(&mailspray.Text, "text", "", "", "TEXT to use")

	sprayCmd.MarkFlagFilename("htmltemplate", "html")
	sprayCmd.MarkFlagFilename("texttemplate", "txt")

	sprayCmd.MarkFlagsMutuallyExclusive("htmltemplate", "texttemplate", "text")

	sprayCmd.Flags().StringVarP(&mailspray.CanaryURL, "canary", "", "", "Canary DNS token to use for Tracking. Will be used for the {{.Canary}} field in a template")

	// sprayCmd.Flags().StringArrayVarP(&mailspray.TemplateVars, "templatevar", "v", []string{}, "Template variables eg: -v Name=Test -v URL=https://Test.com")

	// sprayCmd.Flags().StringArrayVarP(&mailspray.To, "to", "t", []string{}, "Addresses to send email TO.")
	// sprayCmd.MarkFlagRequired("to")
	sprayCmd.Flags().StringVarP(&csvFile, "csv", "", "", "CSV file to parse")
	sprayCmd.MarkFlagFilename("csv", "csv")
	sprayCmd.MarkFlagRequired("csv")

	sprayCmd.PersistentFlags().BoolVar(&MailServer.MaintainConnection, "dont-reuse", false, "Disable the reuse of a single SMTP connection")

	sprayCmd.Flags().StringArrayVarP(&mailspray.HeaderVars, "header", "", []string{}, "Specify additional headers eg: --header Phish=Knowbe4 --header Source=example")

	sprayCmd.Flags().StringArrayVarP(&mailspray.EmbedFiles, "embed", "e", []string{}, "Specify files to embed. These can then be refrenced by their file name in the html (eg: <img src=\"cid:email-logo1.png\">) eg: --embed email-logo1.png --embed email-logo2.png")
	sprayCmd.Flags().StringArrayVarP(&mailspray.AttachFiles, "attach", "a", []string{}, "Specify files to attach. eg: --attach test.pdf --attach average.exe")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sprayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sprayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
