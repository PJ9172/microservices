package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

func SendWelcomeEmail(to string) {
	e := email.NewEmail()
	e.From = "Microservices <" + os.Getenv("SMTP_USER") + ">"
	e.To = []string{to}
	e.Subject = "Welcome to MyApp!"
	e.Text = []byte("Thanks for registering!")

	err := e.Send(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST")),
	)
	if err != nil {
		log.Println("Failed to send email : ", err)
	}
	fmt.Println("Email send successfully!!")

}

func SendPaymentEmail(to string) {
	e := email.NewEmail()
	e.From = "Microservices <" + os.Getenv("SMTP_USER") + ">"
	e.To = []string{to}
	e.Subject = "Payment Status : Done"
	e.Text = []byte("Thanks for Payment!!!")

	err := e.Send(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST")),
	)
	if err != nil {
		log.Println("Failed to send email : ", err)
		return
	}
	fmt.Println("Email send successfully!!")

}
