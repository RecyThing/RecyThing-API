package email

import (
	"io/ioutil"
	"log"
	"strings"
)

func SendOTPEmail(emailAddress string, otp string) {
	go func() {
		// Baca template email dari file.
		emailTemplateBytes, err := ioutil.ReadFile("utils/email/templates/otp.html")
		if err != nil {
			log.Printf("gagal membaca template email: %v", err)
			return
		}
		emailTemplate := string(emailTemplateBytes)

		// Ganti placeholder dalam template dengan OTP sebenarnya.
		emailContent := strings.Replace(emailTemplate, "{{.Otp}}", otp, -1)

		// Kirim email dengan menggunakan fungsi SMTP.
		// Pastikan fungsi SendEmailSMTP memiliki signature yang sesuai dengan penggunaan ini.
		_, errEmail := SendEmailSMTPForOTP([]string{emailAddress}, emailContent, otp)
		if errEmail != nil {
			log.Printf("gagal mengirim otp: %v", errEmail)
		}
	}()
}
