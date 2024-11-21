package main

import (
	"fmt"

	"github.com/chrisjoyce911/usivalidator"
)

func main() {
	usi := "ABCDEF1234" // Example USI

	// Validate the USI
	valid, err := usivalidator.VerifyKey(usi)
	if err != nil {
		fmt.Println("Error:", err)
	} else if valid {
		fmt.Println("The USI is valid!")
	} else {
		fmt.Println("The USI is invalid!")
	}

	// Generate a check character for a USI prefix
	prefix := "ABCDEF123"
	checkChar, err := usivalidator.GenerateCheckCharacter(prefix)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The check character for %s is %c\n", prefix, checkChar)
	}
}
