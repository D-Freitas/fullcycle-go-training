package patternvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAInvalidPhoneNumber_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isMatch := PhoneNumberValidator("00000000000")
	assert.Equal(t, false, isMatch)
}
