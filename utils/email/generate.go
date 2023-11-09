package email

import (
	"crypto/rand"
	"math/big"
	"encoding/base64"
)

func GenerateUniqueToken() string {
	// Generate a unique token
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func GenerateOTP(length int) (string, error) {
	const charset = "1234567890"
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}