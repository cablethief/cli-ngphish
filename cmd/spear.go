/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/cablethief/cli-ngphish/lib"

	"github.com/spf13/cobra"
)

var (
	mail = lib.NewMail()
)

// var templateVars []string

// var htmlFile string
// var textFile string
// var text string

// var subject string

// var from string
// var to string

// var smtpServer string
// var smtpPort int

// var smtpUser string
// var smtpPassword string

// var headerVars []string
// var embedFiles []string
// var attachFiles []string

// spearCmd represents the spear command
var spearCmd = &cobra.Command{
	Use:   "spear",
	Short: "Target a single mailbox",
	Long: `Makes a direct connection to a SMTP server to send an email.
    
    Template may declare substitutions with {{.Name}} syntax.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spear called")
		fmt.Println(mail.smtpServer)
		mail.server = server
		mail.sendMail()
	},
}

func init() {
	rootCmd.AddCommand(spearCmd)

	spearCmd.Flags().StringVarP(&mail.subject, "subject", "", "Good day", "Subject of the email")
	spearCmd.MarkFlagRequired("subject")

	spearCmd.Flags().StringVarP(&mail.htmlFile, "htmltemplate", "", "", "HTML template to use")
	spearCmd.Flags().StringVarP(&mail.textFile, "texttemplate", "", "", "TEXT template to use")
	spearCmd.Flags().StringVarP(&mail.text, "text", "", "", "TEXT to use")

	spearCmd.MarkFlagFilename("htmltemplate", "html")
	spearCmd.MarkFlagFilename("texttemplate", "txt")

	spearCmd.MarkFlagsMutuallyExclusive("htmltemplate", "texttemplate", "text")

	spearCmd.Flags().StringArrayVarP(&mail.templateVars, "templatevar", "v", []string{}, "Template variables eg: -v Name=Test -v URL=https://Test.com")

	spearCmd.Flags().StringVarP(&mail.to, "to", "t", "", "Address to send email TO.")
	spearCmd.MarkFlagRequired("to")
	spearCmd.Flags().StringVarP(&mail.from, "from", "f", "", "Address to send email FROM. eg: Michael <michael@testing.test>")
	spearCmd.MarkFlagRequired("from")

	spearCmd.Flags().StringArrayVarP(&mail.headerVars, "header", "", []string{}, "Specify additional headers eg: --header Phish=Knowbe4 --header Source=example")

	spearCmd.Flags().StringArrayVarP(&mail.embedFiles, "embed", "e", []string{}, "Specify files to embed. These can then be refrenced by their file name in the html (eg: <img src=\"cid:email-logo1.png\">) eg: --embed email-logo1.png --embed email-logo2.png")
	spearCmd.Flags().StringArrayVarP(&mail.attachFiles, "attach", "a", []string{}, "Specify files to attach. eg: --attach test.pdf --attach average.exe")

	// TODO Error handling...
	// TODO BCC CC https://github.com/go-gomail/gomail/issues/19

	spearCmd.Flags().SortFlags = false

}
