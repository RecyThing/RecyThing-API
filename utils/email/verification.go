package email

import (
	"io/ioutil"
	"log"
	"recything/utils/helper"
)

func SendVerificationEmail(emailAddress string, token string) {
	go func() {
		verificationLink := helper.VERIFICATION_URL + token
		emailTemplate, err := ioutil.ReadFile("utils/email/templates/account_registration.html")
		if err != nil {
			log.Printf("gagal membaca template email: %v", err)
			return
		}

		_, errEmail := SendEmailSMTP([]string{emailAddress}, string(emailTemplate), verificationLink)
		if errEmail != nil {
			log.Printf("gagal mengirim email verifikasi: %v", errEmail)
		}
	}()
}
