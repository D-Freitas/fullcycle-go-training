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

func TestGivenAPasswordWithoutUppercaseLetters_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isValid := PasswordValidator("davifreitas@123")
	assert.Equal(t, false, isValid)
}

func TestGivenAPasswordWithUppercaseLetters_WhenToCheck_ThenMatchShouldReturnTrue(t *testing.T) {
	isValid := PasswordValidator("DaviFreitas@123")
	assert.Equal(t, true, isValid)
}

func TestGivenAPasswordWithoutSpecialChars_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isValid := PasswordValidator("DaviFreitas123")
	assert.Equal(t, false, isValid)
}

func TestGivenAPasswordWithSpecialChars_WhenToCheck_ThenMatchShouldReturnFalse(t *testing.T) {
	isValid := PasswordValidator("Davi@Freitas123")
	assert.Equal(t, false, isValid)
}
