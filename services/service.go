package services

import sendEmail "NaimBiswas/email-reminder/services/send-email"

func SendEmailService(to []string, hrNames string) {
	email := sendEmail.New(to, hrNames)
	email.SendEmailToClient()
}
