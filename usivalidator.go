/*
Package usivalidator provides tools to validate and generate the check character for Unique Student Identifiers (USI) in Australia.

What is a Unique Student Identifier (USI)?
A USI is your individual education number for life. It also provides you with an online record of your vocational education and training (VET) undertaken in Australia.

If you're at university, TAFE, or engaged in other nationally recognised training, you need a USI. Without one, you cannot receive Commonwealth financial assistance or obtain your qualification or statement of attainment.

This package implements the validation process using the Luhn Mod N algorithm, which calculates the check character based on a specific set of valid characters.
*/
package usivalidator

import (
	"errors"
	"strings"
)

// ValidCharacters contains the valid characters for the USI
var ValidCharacters = []rune{'2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R',
	'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// VerifyKey validates a 10-character USI against its calculated check character.
//
// Parameters:
// - key (string): The USI to validate. Must be exactly 10 characters long.
//
// Returns:
// - (bool): True if the USI is valid, false otherwise.
// - (error): An error if the input length is invalid or contains invalid characters.
//
// Usage:
// isValid, err := VerifyKey("BNGH7C75FN")
// if err != nil {
//     log.Println("Error:", err)
// } else if isValid {
//     fmt.Println("The USI is valid!")
// } else {
//     fmt.Println("The USI is invalid!")
// }

func VerifyKey(key string) (bool, error) {
	if len(key) != 10 {
		return false, errors.New("key length must be 10 characters")
	}

	key = strings.ToUpper(key)
	checkDigit, err := GenerateCheckCharacter(key[:9])
	if err != nil {
		return false, err
	}

	return rune(key[9]) == checkDigit, nil
}

// GenerateCheckCharacter calculates the check character for a 9-character USI prefix
// using the Luhn Mod N algorithm.
//
// Parameters:
// - input (string): The first 9 characters of the USI.
//
// Returns:
// - (rune): The calculated check character.
// - (error): An error if the input length is not 9 characters or contains invalid characters.
//
// Usage:
// checkChar, err := GenerateCheckCharacter("BNGH7C75F")
// if err != nil {
//     log.Println("Error:", err)
// } else {
//     fmt.Printf("The check character is %c\n", checkChar)
// }

func GenerateCheckCharacter(input string) (rune, error) {
	if len(input) != 9 {
		return ' ', errors.New("input length must be 9 characters")
	}

	factor := 2
	sum := 0
	n := len(ValidCharacters)

	for i := len(input) - 1; i >= 0; i-- {
		char := rune(input[i])
		codePoint := indexOf(char, ValidCharacters)
		if codePoint == -1 {
			return ' ', errors.New("invalid character in input")
		}

		addend := factor * codePoint
		factor = alternateFactor(factor)
		addend = (addend / n) + (addend % n)
		sum += addend
	}

	remainder := sum % n
	checkCodePoint := (n - remainder) % n

	return ValidCharacters[checkCodePoint], nil
}

// indexOf finds the index of a rune in a slice of runes.
//
// Parameters:
// - char (rune): The character to find.
// - slice ([]rune): The slice of valid characters.
//
// Returns:
// - (int): The index of the character in the slice, or -1 if not found.
//
// Usage:
// index := indexOf('A', ValidCharacters)
// if index != -1 {
//     fmt.Printf("Character found at index %d\n", index)
// } else {
//     fmt.Println("Character not found")
// }

func indexOf(char rune, slice []rune) int {
	for i, v := range slice {
		if v == char {
			return i
		}
	}
	return -1
}

// alternateFactor alternates between the multiplication factors used in the Luhn Mod N algorithm.
//
// Parameters:
// - factor (int): The current factor, either 1 or 2.
//
// Returns:
// - (int): The alternate factor (2 if the input is 1, or 1 if the input is 2).
//
// Usage:
// nextFactor := alternateFactor(2) // Returns 1

func alternateFactor(factor int) int {
	if factor == 2 {
		return 1
	}
	return 2
}
