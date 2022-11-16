package patternvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAPasswordWithoutSevenLetters_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isValid := PasswordValidator("davi@123")
	assert.Equal(t, false, isValid)
}

func TestGivenAPasswordWithSevenOrMoreLetters_WhenToCheck_ThenMatchShouldReturnTrue(t *testing.T) {
	isValid := PasswordValidator("DaviFreitas@123")
	assert.Equal(t, true, isValid)
}