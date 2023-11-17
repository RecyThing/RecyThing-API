package validation

import (
	"errors"
	"recything/utils/constanta"
	"regexp"
	"strconv"
	"strings"
)

func CheckDataEmpty(data ...any) error {
	for _, value := range data {
		if value == "" {
			return errors.New(constanta.ERROR_EMPTY)
		}
	}
	return nil
}


func CheckCategory(data string, validCategories []string) (string, error) {
	inputCategory := strings.ToLower(data)

	isValidCategory := false
	for _, category := range validCategories {
		if inputCategory == strings.ToLower(category) {
			isValidCategory = true
			break
		}
	}

	if !isValidCategory {
		return "", errors.New("data yang diinput tidak sesuai")
	}

	return inputCategory, nil
}

func EmailFormat(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return nil
	}
	return errors.New(constanta.ERROR_FORMAT_EMAIL)
}

func PhoneNumber(phone string) error {
	if len(phone) < 10 || len(phone) > 16 {
		return errors.New("panjang nomor telepon harus antara 10 dan 16 karakter")
	}

	phoneRegex := `^(0811|0812|0813|0821|0822|0823|0851|0852|0853|0814|0815|0816|0855|0856|0857|0858|0895|0896|0897|0898|0899|0817|0818|0819|0859|0877|0878|0879|0881|0882|0883|0885|0886|0887|0888|0889|0810|0854|0880|0884|0889|0891|0892|0893|0894|0896|0897|0899|62811|62812|62813|62821|62822|62823|62851|62852|62853|62814|62815|62816|62855|62856|62857|62858|62895|62896|62897|62898|62899|62817|62818|62819|62859|62877|62878|62879|62881|62882|62883|62885|62886|62887|62888|62889|62810|62854|62880|62884|62889|62891|62892|62893|62894|62896|62897|62899)\d{8}$`
	regex := regexp.MustCompile(phoneRegex)

	if regex.MatchString(phone) {
		return nil
	}

	return errors.New("format nomor telepon tidak valid")
}

func MinLength(data string, minLength int) error {
	if len(data) < minLength {
		return errors.New("minimal " + strconv.Itoa(minLength) + " karakter,ulangi kembali!")
	}
	return nil
}
