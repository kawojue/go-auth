package utils

import (
	"math/rand"
	"time"

	"github.com/kawojue/go-auth/structs"
)

func GenOTP(length int) structs.TOTP {
	var otp string
	digits := "0123456789"

	for i := 0; i < length; i++ {
		otp += string(digits[rand.Intn(len(digits))])
	}

	otp_expiry := time.Now().Add(
		30 * time.Minute,
	).UTC().Format(
		"2006-01-02T15:04:05.999Z",
	)

	return structs.TOTP{
		Otp:        otp,
		Otp_Expiry: otp_expiry,
	}
}
