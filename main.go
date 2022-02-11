package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// Configuration
	from := "golang7744@gmail.com"
	password := "g123456789?"
	to := []string{
		"azizxonkom@gmail.com",
		"qodircodermmi@gmail.com",
		"komronfayziboev@gmail.com",
		"xurshidbeksobirov03@gamil.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Ishladi")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Supper")

}
