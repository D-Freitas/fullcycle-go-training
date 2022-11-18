package patternvalidator

import "regexp"

func PhoneNumberValidator(phoneNumber string) bool {
	isFilled := phoneNumber != ""
	phoneNumberRegexp := regexp.MustCompile(`^\d{10,11}$`)
	isMatch := phoneNumberRegexp.MatchString(phoneNumber)
	return isFilled && isMatch
}
