package patternvalidator

import (
	"unicode"
)

func PasswordValidator(password string) bool {
	letters := 0
	for _, c := range password {
		switch {
		case unicode.IsLetter(c) || c == ' ':
			letters++
		}
	}
	containsSevenLettersOrMore := letters >= 7
	isValid := containsSevenLettersOrMore
	return isValid
}
