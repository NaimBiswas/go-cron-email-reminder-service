package sendEmail

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strconv"
	"text/template"
	"time"
)

type SendEmailStruct struct{}

var (
	body          bytes.Buffer
	smtpAuth      smtp.Auth
	toEmail       []string
	hrManagerName string
	smtpPort      string = "587"
	smtpServer           = "smtp.gmail.com"
	smtpUser             = "naimbiswas0"
	smtpPassword         = "jekasojgnnufgpcl"
)

func authenticateSmtp() {
	smtpAuth = smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)
}

func New(to []string, hrNames string) *SendEmailStruct {
	toEmail = to
	hrManagerName = hrNames
	authenticateSmtp()
	prepareEmailBody()

	return &SendEmailStruct{}

}

func (s *SendEmailStruct) SendEmailToClient() {

	now := time.Now()
	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastMonth := firstDayOfMonth.AddDate(0, -1, 0)
	previousMonth := lastMonth.Format("January 2006")

	timestamp := now.Unix()

	from := "Naim Biswas"
	subject := "Friendly Reminder: Pending Salary for " + previousMonth + " (" + strconv.FormatInt(timestamp, 10) + ")"
	msg := []byte("From: " + from + "\r\n" +
		"To: " + fmt.Sprintf("%s", toEmail) + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body.String())

	err := smtp.SendMail(smtpServer+":"+smtpPort, smtpAuth, from, toEmail, msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Email sent successfully")
}

func prepareEmailBody() {
	now := time.Now()

	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastMonth := firstDayOfMonth.AddDate(0, -1, 0)
	previousMonth := lastMonth.Format("January 2006")

	// Read the HTML template from a file
	templateFile := "templates/salary-reminder.template.html"
	tmpltContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fmt.Println(err)
	}

	// Create the email content template
	tmpl, err := template.New("email").Parse(string(tmpltContent))
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		PreviousMonth string
		HRManagerName string
	}{
		PreviousMonth: previousMonth,
		HRManagerName: hrManagerName,
	}
	var emailBody bytes.Buffer
	err = tmpl.Execute(&emailBody, data)

	if err != nil {
		fmt.Println(err)
	}
	body = emailBody
}
