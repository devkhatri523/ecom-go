package service

import (
	"bytes"
	"fmt"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	"html/template"
	"net/http"
	"net/smtp"
	"strconv"
	"v01/domain"
)

type OrderConfirmationEmailSender struct {
}

func NewOrderConfirmationEmailSender() EmailService {
	return &OrderConfirmationEmailSender{}
}

func (o OrderConfirmationEmailSender) SendEmail(v interface{}) error {
	orderConfirmation := v.(domain.OrderConfirmation)
	auth := smtp.PlainAuth("", config.Default().GetString("mail.username"), config.Default().GetString("mail.password"),
		config.Default().GetString("mail.host"))
	tmpl, err := template.ParseFiles("email_template.html")
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return utils.NewError(strconv.Itoa(http.StatusInternalServerError), err)
	}
	var body bytes.Buffer
	err = tmpl.Execute(&body, orderConfirmation)
	if err != nil {
		return err
	}

	msg := []byte("To: recipient@example.com\r\n" +
		"Subject: Order Confirmation email from ecommerce site\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		body.String())
	err = smtp.SendMail(
		"localhost:1025",
		auth,
		"noreply@gmail.com",
		nil,
		msg,
	)

	if err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	} else {
		fmt.Println("Email sent successfully!")
	}
	return nil
}
