package patternvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnInvalidEmail_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isMatch := EmailValidator("invalid_email")
	assert.Equal(t, false, isMatch)
}

func TestGivenAnValidEmail_WhenToCheck_ThenMatchShouldReturnTrue(t *testing.T) {
	isMatch := EmailValidator("davi.israel01@gmail.com")
	assert.Equal(t, true, isMatch)
}

func TestGivenAnEmptyEmail_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isMatch := EmailValidator("")
	assert.Equal(t, false, isMatch)
}
