package patternvalidator

import "regexp"

func EmailValidator(email string) bool {
	isFilled := email != ""
	emailRegexp := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	isMatch := emailRegexp.MatchString(email)
	return isFilled && isMatch
}
