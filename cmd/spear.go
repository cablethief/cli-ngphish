/*
Copyright © 2022 Michael Kruger @_Cablethief
*/
package cmd

import (
	"log"

	"github.com/cablethief/cli-ngphish/lib"

	"github.com/spf13/cobra"
)

var mail = &lib.Mail{}

// spearCmd represents the spear command
var spearCmd = &cobra.Command{
	Use:   "spear",
	Short: "Target a single mailbox",
	Long: `Targets a single mailbox using the specified template or text:
    
    Template may declare substitutions with {{.VarName}} syntax.`,
	Run: func(cmd *cobra.Command, args []string) {

		MailServer.SendMail(mail.Create())
		log.Println(mail)
	},
}

func init() {
	rootCmd.AddCommand(spearCmd)

	spearCmd.Flags().StringVarP(&mail.To, "to", "t", "", "Addresses to send email TO.")
	spearCmd.MarkFlagRequired("to")

	spearCmd.Flags().StringVarP(&mail.From, "from", "f", "", "Address to send email FROM. eg: Michael <michael@testing.test>")
	spearCmd.MarkFlagRequired("from")

	spearCmd.Flags().StringArrayVarP(&mail.Cc, "cc", "", []string{}, "Address to CC in the email.")
	spearCmd.Flags().StringArrayVarP(&mail.Bcc, "bcc", "", []string{}, "Addresses to BCC in the email.")

	spearCmd.Flags().StringVarP(&mail.Subject, "subject", "", "Good day", "Subject of the email")
	spearCmd.MarkFlagRequired("subject")

	spearCmd.Flags().StringVarP(&mail.HtmlFile, "htmltemplate", "", "", "HTML template to use")
	spearCmd.Flags().StringVarP(&mail.TextFile, "texttemplate", "", "", "TEXT template to use")
	spearCmd.Flags().StringVarP(&mail.Text, "text", "", "", "TEXT to use")

	spearCmd.MarkFlagFilename("htmltemplate", "html")
	spearCmd.MarkFlagFilename("texttemplate", "txt")

	spearCmd.MarkFlagsMutuallyExclusive("htmltemplate", "texttemplate", "text")

	spearCmd.Flags().StringArrayVarP(&mail.TemplateVars, "templatevar", "v", []string{}, "Template variables eg: -v Name=Test -v URL=https://Test.com")

	spearCmd.Flags().StringVarP(&mail.CanaryURL, "canary", "", "", "Canary DNS token to use for Tracking. Will be used for the {{.Canary}} field in a template")

	spearCmd.Flags().StringArrayVarP(&mail.HeaderVars, "header", "", []string{}, "Specify additional headers eg: --header Phish=Knowbe4 --header Source=example")

	spearCmd.Flags().StringArrayVarP(&mail.EmbedFiles, "embed", "e", []string{}, "Specify files to embed. These can then be refrenced by their file name in the html (eg: <img src=\"cid:email-logo1.png\">) eg: --embed email-logo1.png --embed email-logo2.png")
	spearCmd.Flags().StringArrayVarP(&mail.AttachFiles, "attach", "a", []string{}, "Specify files to attach. eg: --attach test.pdf --attach average.exe")

	// TODO Error handling...
	// TODO BCC CC https://github.com/go-gomail/gomail/issues/19

	spearCmd.Flags().SortFlags = false

}
