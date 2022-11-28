package validator

import "regexp"

func PhoneNumberValidator(phoneNumber string) bool {
	phoneNumberRegexp := regexp.MustCompile(`^\d{10,11}$`)
	isMatch := phoneNumberRegexp.MatchString(phoneNumber)
	return isMatch
}
