package clients

import (
	"net/mail"
	"strconv"
)

func validE(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validPhone(phone string) bool {
	if _, err := strconv.ParseInt(phone, 10, 64); err == nil {
		return len(phone) == 9
	} else {
		return false
	}
}
