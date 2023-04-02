package service

import (
	"cbupnvj/config"
	"context"
	"net/smtp"
	"strings"

	"github.com/sirupsen/logrus"
)

func sendMail(ctx context.Context, toMail []string, subject, body string) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"toMail": toMail,
	})
	// Set up authentication information.
	auth := smtp.PlainAuth("", config.MailUsername(), config.MailPassword(), "smtp.gmail.com")

	// Set up the message.
	message := "From: " + config.MailUsername() + "\n" +
		"To: " + strings.Join(toMail, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	// Send the message.
	err := smtp.SendMail("smtp.gmail.com:587", auth, config.MailUsername(), toMail, []byte(message))
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Success send mail")
	}
}
