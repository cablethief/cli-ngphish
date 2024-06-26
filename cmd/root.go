/*
Copyright © 2022 Michael Kruger @_Cablethief
*/
package cmd

import (
	"os"

	"github.com/cablethief/cli-ngphish/lib"
	"github.com/spf13/cobra"
)

var MailServer = lib.NewMailServer()

// var mailserver = lib.NewMailServer()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-ngphish",
	Short: "A CLI tool to assist with Phishing",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mailspoofcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&MailServer.SmtpServer, "server", "s", "", "SMTP server to make a connection to. eg: tenant.mail.protection.outlook.com")
	spearCmd.MarkFlagRequired("server")
	rootCmd.PersistentFlags().IntVarP(&MailServer.SmtpPort, "port", "p", 25, "SMTP server port to make a connection to")
	rootCmd.PersistentFlags().BoolVar(&MailServer.CheckTLS, "tls", false, "Whether to validate the certificate presented by the SMTP server")

	rootCmd.PersistentFlags().StringVar(&MailServer.SmtpUser, "username", "", "SMTP Username")
	rootCmd.PersistentFlags().StringVar(&MailServer.SmtpPassword, "password", "", "SMTP Password")
	rootCmd.MarkFlagsRequiredTogether("username", "password")
}
