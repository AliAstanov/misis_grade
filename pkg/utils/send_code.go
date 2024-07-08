package utils

import (
	"fmt"
	"log"
	"net/smtp"
)

func Otp() {
	otp, err := GenerateOTP(6)
	if err != nil {
		log.Printf("Failed to generate OTP: %v", err)
	}

	to := []string{"fayzulloyevismoil9@gmail.com","aliastan1997@gmail.com"}
	err = SendEmail(to, otp)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
	}

	fmt.Printf("OTP code %s sent to %s\n", otp, to)
}

func SendEmail(to []string, otp string) error {

	var (
		fromGmail = "alibekastan1998@gmail.com"
		password  = "jeed qqdo stze gtgm "	
	)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: Your OTP Code\n\n" + "Your OTP code is: " + otp)

	bytM := []byte(message)

	auth := smtp.PlainAuth("Alibek_!!!", fromGmail, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromGmail, to, bytM)
	if err != nil {
		return err
	}
	return nil
}
