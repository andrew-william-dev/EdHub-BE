package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmailOTP(recipient, otp string) error {
	mailKey := os.Getenv("MAIL_KEY")
	from := "edhub2025@gmail.com"
	password := mailKey
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Subject: EdHub - OTP Verification\n"
	body := fmt.Sprintf(`
	<html>
		<body style="font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f4f7fb; padding: 20px; margin: 0;">
			<div style="max-width: 600px; margin: 0 auto; padding: 25px; background-color: #ffffff; border-radius: 12px; box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);">
				<h2 style="color: #333333; text-align: center; font-size: 26px; font-weight: 600;">EdHub OTP Verification</h2>
				<p style="font-size: 16px; color: #666666; text-align: center; margin-top: 10px;">Hello,</p>
				<p style="font-size: 16px; color: #666666; text-align: center;">Your OTP (One-Time Password) for verification is:</p>
				<h3 style="font-size: 32px; color: #ffffff; text-align: center; padding: 20px; background: linear-gradient(135deg, #ff6f61, #ff3d00); border-radius: 12px; box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1); letter-spacing: 3px; font-weight: bold; margin-top: 20px; margin-bottom: 20px;">
					%s
				</h3>
				<p style="font-size: 16px; color: #666666; text-align: center; margin-bottom: 20px;">Please use this code to complete your registration.</p>
				<hr style="border: 0; border-top: 1px solid #ddd; margin-bottom: 20px;">
				<p style="font-size: 14px; color: #999999; text-align: center;">If you did not request this code, please ignore this email.</p>
			</div>
		</body>
	</html>`, otp)

	message := []byte("MIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n" + subject + "\r\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipient}, message)
	return err
}

func SendForgotMailOTP(recipient, otp string) error {
	mailKey := os.Getenv("MAIL_KEY")
	from := "edhub2025@gmail.com"
	password := mailKey
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Subject: EdHub Register New\n"
	body := fmt.Sprintf("Your OTP for verification is: %s", otp)
	message := []byte(subject + "\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipient}, message)
	return err
}
