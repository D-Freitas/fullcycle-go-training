package patternvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnInvalidEmail_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isMatch := EmailValidator("invalid_email")
	assert.Equal(t, false, isMatch)
}
