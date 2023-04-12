/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// sprayCmd represents the spray command
var sprayCmd = &cobra.Command{
    Use:   "spray",
    Short: "Spray at multiple mailboxes based off a CSV",
    Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Run: func(cmd *cobra.Command, args []string) {



        
        fmt.Println("spray called")



    },
}

func init() {
    rootCmd.AddCommand(sprayCmd)

    sprayCmd.Flags().StringVarP(&smtpServer, "server", "s", "", "SMTP server to make a connection to. eg: tenant.mail.protection.outlook.com")
    sprayCmd.MarkFlagRequired("server")
    sprayCmd.Flags().IntVarP(&smtpPort, "port", "p", 25, "SMTP server port to make a connection to")

    sprayCmd.Flags().StringVar(&smtpUser, "username", "", "SMTP Username")
    sprayCmd.Flags().StringVar(&smtpPassword, "password", "", "SMTP Password")
    sprayCmd.MarkFlagsRequiredTogether("username", "password")


    sprayCmd.Flags().StringVarP(&subject, "subject", "", "Good day", "Subject of the email")
    sprayCmd.MarkFlagRequired("subject")


    sprayCmd.Flags().StringVarP(&htmlFile, "htmltemplate", "", "", "HTML template to use")
    sprayCmd.Flags().StringVarP(&textFile, "texttemplate", "", "", "TEXT template to use")
    sprayCmd.Flags().StringVarP(&text, "text", "", "", "TEXT to use")

    sprayCmd.MarkFlagFilename("htmltemplate","html")
    sprayCmd.MarkFlagFilename("texttemplate","txt")

    sprayCmd.MarkFlagsMutuallyExclusive("htmltemplate", "texttemplate", "text")

    sprayCmd.Flags().StringArrayVarP(&headerVars, "header", "", []string{}, "Specify additional headers eg: --header Phish=Knowbe4 --header Source=example")

    sprayCmd.Flags().StringArrayVarP(&embedFiles, "embed", "e", []string{}, "Specify files to embed. These can then be refrenced by their file name in the html (eg: <img src=\"cid:email-logo1.png\">) eg: --embed email-logo1.png --embed email-logo2.png")
    sprayCmd.Flags().StringArrayVarP(&attachFiles, "attach", "a", []string{}, "Specify files to attach. eg: --attach test.pdf --attach average.exe")




    sprayCmd.Flags().StringArrayVarP(&templateVars, "templatevar", "v", []string{}, "Template variables eg: -v Name=Test -v URL=https://Test.com")

    sprayCmd.Flags().StringVarP(&to, "to", "t", "", "Address to send email TO.")
    sprayCmd.MarkFlagRequired("to")
    sprayCmd.Flags().StringVarP(&from, "from", "f", "", "Address to send email FROM. eg: Michael <michael@testing.test>")
    sprayCmd.MarkFlagRequired("from")




    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // sprayCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // sprayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
