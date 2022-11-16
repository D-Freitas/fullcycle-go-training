package patternvalidator

import (
	"unicode"
)

func PasswordValidator(password string) bool {
	letters := 0
	isUpper := false
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			isUpper = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		}
	}
	containsSevenLettersOrMore := letters >= 7
	isValid := containsSevenLettersOrMore && isUpper
	return isValid
}
