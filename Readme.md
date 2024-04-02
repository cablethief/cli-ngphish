# Cli-ngphish 📧

A CLI tool for sending phishing mails with templating options. I felt that in a lot of cases I didn't want to setup a whole gophish instance and would rather manage some of the stuff myself. 

The tool is essentially a nice wrapper around the lovely [gomail](https://github.com/go-gomail/gomail) lib in order to template mails for both spear phishing and spraying. 

I have mostly made use of this tool to test spoofing against office365 as it is quick to play around with different options and from addresses without the clutter of gophishes interface. 

```
Usage:
  cli-ngphish [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  spear       Target a single mailbox
  spray       Spray at multiple mailboxes based off a CSV
  version     A brief description of your command

Flags:
  -h, --help              help for cli-ngphish
      --password string   SMTP Password
  -p, --port int          SMTP server port to make a connection to (default 25)
  -s, --server string     SMTP server to make a connection to. eg: tenant.mail.protection.outlook.com
      --tls               Whether to validate the certificate presented by the SMTP server
      --username string   SMTP Username
```

# Spear phishing

```
Usage:
  cli-ngphish spear [flags]

Flags:
  -t, --to string                 Addresses to send email TO.
  -f, --from string               Address to send email FROM. eg: Michael <michael@testing.test>
      --cc stringArray            Address to CC in the email.
      --bcc stringArray           Addresses to BCC in the email.
      --subject string            Subject of the email (default "Good day")
      --htmltemplate string       HTML template to use
      --texttemplate string       TEXT template to use
      --text string               TEXT to use
  -v, --templatevar stringArray   Template variables eg: -v Name=Test -v URL=https://Test.com
      --canary string             Canary DNS token to use for Tracking. Will be used for the {{.Canary}} field in a template
      --header stringArray        Specify additional headers eg: --header Phish=Knowbe4 --header Source=example
  -e, --embed stringArray         Specify files to embed. These can then be refrenced by their file name in the html (eg: <img src="cid:email-logo1.png">) eg: --embed email-logo1.png --embed email-logo2.png
  -a, --attach stringArray        Specify files to attach. eg: --attach test.pdf --attach average.exe
  -h, --help                      help for spear
```

# Spray

The spray command will want a csv file which has the field `to`. Additional fields can be added which will be used in the template and can be referenced with `{{.Name}}`. This allows for long lists of users with custom variables per user.  

```
to, Name, Surname
michael.kruger@orangecyberdefense.com, Michael, Kruger
felipe.molinadelatorre@orangecyberdefense.com, Felipe, Kruger2
szymon.ziolkowski@orangecyberdefense.com, Szymon, ASJKDHJKSAND

```
Usage:
  cli-ngphish spray [flags]

Flags:
  -a, --attach stringArray    Specify files to attach. eg: --attach test.pdf --attach average.exe
      --canary string         Canary DNS token to use for Tracking. Will be used for the {{.Canary}} field in a template
      --csv string            CSV file to parse
      --dont-reuse            Disable the reuse of a single SMTP connection
  -e, --embed stringArray     Specify files to embed. These can then be refrenced by their file name in the html (eg: <img src="cid:email-logo1.png">) eg: --embed email-logo1.png --embed email-logo2.png
  -f, --from string           Address to send email FROM. eg: Michael <michael@testing.test>
      --header stringArray    Specify additional headers eg: --header Phish=Knowbe4 --header Source=example
  -h, --help                  help for spray
      --htmltemplate string   HTML template to use
      --subject string        Subject of the email (default "Good day")
      --text string           TEXT to use
      --texttemplate string   TEXT template to use
```