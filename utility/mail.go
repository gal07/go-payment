package utility

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Student App <galihkur29@gmail.com>"
const CONFIG_AUTH_EMAIL = "galihkur29@gmail.com"
const CONFIG_AUTH_PASSWORD = "aqogepbxmthlljpl"

func Send(target []string, target_cc []string, activation string) {

	to := target
	cc := target_cc
	subject := "Activation user"
	data := activation
	message := "here to activation user <a href='http://localhost:8041/api/v1/auth/gotoactive?activation='" + data + "></a>"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		CreateLog(err)
		// log.Fatal(err.Error())
	}

	log.Println("Mail sent!")

}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		CreateLog(err)
		return err
	}

	return nil
}
