package usivalidator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleVerifyKey() {
	usi := "BNGH7C75FN" // Example USI
	isValid, err := VerifyKey(usi)
	if err != nil {
		fmt.Println("Error:", err)
	} else if isValid {
		fmt.Println("The USI is valid!")
	} else {
		fmt.Println("The USI is invalid!")
	}

	// Output: The USI is valid!
}

func TestVerifyKey(t *testing.T) {
	testCases := []struct {
		USI         string
		IsValid     bool
		ExpectedErr string
	}{
		{"BNGH7C75FN", true, ""},                                 // Valid USI
		{"BP6LKB3C7X", true, ""},                                 // Valid USI
		{"RVJ5DM8LXJ", true, ""},                                 // Valid USI
		{"PDGGW5XLXW", true, ""},                                 // Valid USI
		{"DG6K5YHPP3", true, ""},                                 // Valid USI
		{"U6Q8JN6UD9", true, ""},                                 // Valid USI
		{"R5HQLSWS9", false, "key length must be 10 characters"}, // Invalid length
		{"INVALID!X", false, "key length must be 10 characters"}, // Invalid character
		{"ABCDEF123@", false, "invalid character in input"},      // Invalid special character
		{"", false, "key length must be 10 characters"},          // Empty string
	}

	for _, tc := range testCases {
		t.Run(tc.USI, func(t *testing.T) {
			isValid, err := VerifyKey(tc.USI)
			if tc.ExpectedErr != "" {
				// Expecting an error
				assert.False(t, isValid, "Expected validation to fail")
				assert.Error(t, err, "Expected an error")
				assert.EqualError(t, err, tc.ExpectedErr, "Error message mismatch")
			} else {
				// No error expected
				assert.True(t, isValid, "Expected validation to succeed")
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}

func ExampleGenerateCheckCharacter() {
	prefix := "BNGH7C75F" // Example 9-character prefix
	checkChar, err := GenerateCheckCharacter(prefix)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The check character is %c\n", checkChar)
	}

	// Output: The check character is N
}

func TestGenerateCheckCharacter(t *testing.T) {
	testCases := []struct {
		Input          string
		ExpectedOutput rune
		ExpectedError  string
	}{
		{"BNGH7C75F", 'N', ""},
		{"BP6LKB3C7", 'X', ""},
		{"RVJ5DM8LX", 'J', ""},
		{"PDGGW5XLX", 'W', ""},
		{"DG6K5YHPP", '3', ""},
		{"U6Q8JN6UD", '9', ""},
		{"INVALIDIN", ' ', "invalid character in input"},
		{"TOOSHORT", ' ', "input length must be 9 characters"},
		{"TOOLONGINPUT", ' ', "input length must be 9 characters"},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			output, err := GenerateCheckCharacter(tc.Input)
			if tc.ExpectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.ExpectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedOutput, output)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	validChars := []rune{'2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	testCases := []struct {
		Char     rune
		Slice    []rune
		Expected int
		TestName string
	}{
		{'A', validChars, 8, "Character found at index 8"},
		{'9', validChars, 7, "Character found at index 7"},
		{'Z', validChars, 31, "Character found at index 31"},
		{'X', validChars, 29, "Character found at index 29"},
		{'1', validChars, -1, "Character not found"},
		{'$', validChars, -1, "Special character not found"},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			result := indexOf(tc.Char, tc.Slice)
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func TestAlternateFactor(t *testing.T) {
	testCases := []struct {
		Input    int
		Expected int
		TestName string
	}{
		{2, 1, "Switch from 2 to 1"},
		{1, 2, "Switch from 1 to 2"},
		{0, 2, "Default case for invalid input (0)"},
		{-1, 2, "Default case for negative input"},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			result := alternateFactor(tc.Input)
			assert.Equal(t, tc.Expected, result)
		})
	}
}
