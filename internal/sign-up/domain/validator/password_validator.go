package validator

import (
	"unicode"
)

func PasswordValidator(password string) bool {
	letters := 0
	isUpper, isSymbol, isNumber := false, false, false
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			isUpper = true
			letters++
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			isSymbol = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		case unicode.IsNumber(c):
			isNumber = true
		}
	}
	isFilled := password != ""
	containsSevenLettersOrMore := letters >= 7
	isValid := isFilled && containsSevenLettersOrMore && isUpper && isSymbol && isNumber
	return isValid
}
