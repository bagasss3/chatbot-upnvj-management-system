package service

import (
	"cbupnvj/config"
	"context"
	"net/smtp"

	"github.com/sirupsen/logrus"
)

func sendMail(ctx context.Context, toMail, subject, body string) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"toMail": toMail,
	})
	// Set up authentication information.
	auth := smtp.PlainAuth("", config.MailUsername(), config.MailPassword(), "smtp.gmail.com")

	// Set up the message.
	to := []string{toMail}
	message := []byte("From: " + config.MailUsername() + "\n" +
		"To: " + toMail + "\n" +
		"Subject: " + subject + "\n\n" +
		body)

	// Send the message.
	err := smtp.SendMail("smtp.gmail.com:587", auth, config.MailUsername(), to, message)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Success send mail")
	}
}
