package patternvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAPasswordWithoutSevenLetters_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isValid := PasswordValidator("davi@123")
	assert.Equal(t, false, isValid)
}
