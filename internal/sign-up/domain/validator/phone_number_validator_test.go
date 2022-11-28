package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAInvalidPhoneNumber_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isMatch := PhoneNumberValidator("000000000")
	assert.Equal(t, false, isMatch)
}

func TestGivenAValidPhoneNumber_WhenToCheck_ThenMatchShouldReturnTrue(t *testing.T) {
	isMatch := PhoneNumberValidator("11925550173")
	assert.Equal(t, true, isMatch)
}
