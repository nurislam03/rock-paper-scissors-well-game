package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testTemplate(t *testing.T, testName, input string, result bool) {
	t.Run("pass", func(t *testing.T) {
		assert.Equal(t, true, validateGameCharacter("rock"))
	})
}

func TestValidateGameCharacture(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		result bool
	}{
		{"for input 'rock' it should return true", "rock", true},
	}

	for _, tc := range testCases {
		testTemplate(t, tc.name, tc.input, tc.result)
	}
}
