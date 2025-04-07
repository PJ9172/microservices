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
	fmt.Println(e.From)
	fmt.Println(e.To)
	fmt.Println(e.Subject)
	fmt.Println(e.Text)
	fmt.Println(os.Getenv("SMTP_USER"))
	fmt.Println(os.Getenv("SMTP_HOST"))
	fmt.Println(os.Getenv("SMTP_PORT"))
	fmt.Println(os.Getenv("SMTP_PASS"))

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
	e.Text = []byte("Thanks for Purchasing Products!!!")

	err := e.Send(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST")),
	)
	if err != nil {
		log.Println("Failed to send email : ", err)
	}
	fmt.Println("Email send successfully!!")

}
