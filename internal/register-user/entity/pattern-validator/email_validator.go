package patternvalidator

import "regexp"

func EmailValidator(email string) bool {
	emailRegexp := regexp.MustCompile("/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$/")
	isMatch := emailRegexp.MatchString(email)
	return isMatch
}
