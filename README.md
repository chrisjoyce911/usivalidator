# USI Validator

## Overview

The `usivalidator` package provides tools to validate and generate check characters for **Unique Student Identifiers (USI)** in Australia. It implements the **Luhn Mod N algorithm** to calculate the check character for a USI and validate its correctness.

### What is a USI?

A **USI** is an individual education number for life. It allows you to maintain an online record of your vocational education and training (VET) undertaken in Australia. If you're at university, TAFE, or undertaking nationally recognized training, you need a USI to receive Commonwealth financial assistance or obtain your qualifications.

---

## Features

- **Validate a USI**: Checks if a given 10-character USI is valid.
- **Generate a Check Character**: Calculates the check character for a given 9-character prefix using the **Luhn Mod N algorithm**.
- Utility functions:
  - Find the index of a character in the valid character set.
  - Alternate factors used in the Luhn Mod N algorithm.

---

## Installation

Install the package using `go get`:

```bash
go get github.com/chrisjoyce911/usivalidator
```

## Example

### Validate a USI

```go
package main

import (
	"fmt"
	"log"

	"github.com/chrisjoyce911/usivalidator"
)

func main() {
	usi := "BNGH7C75FN" // Example USI
	isValid, err := usivalidator.VerifyKey(usi)
	if err != nil {
		log.Println("Error:", err)
	} else if isValid {
		fmt.Println("The USI is valid!")
	} else {
		fmt.Println("The USI is invalid!")
	}
}
```

### Generate a Check Character

```go
package main

import (
	"fmt"
	"log"

	"github.com/chrisjoyce911/usivalidator"
)

func main() {
	usi := "BNGH7C75FN" // Example USI
	isValid, err := usivalidator.VerifyKey(usi)
	if err != nil {
		log.Println("Error:", err)
	} else if isValid {
		fmt.Println("The USI is valid!")
	} else {
		fmt.Println("The USI is invalid!")
	}
}
```

### Utility Functions

```go
package main

import (
	"fmt"
	"log"

	"github.com/chrisjoyce911/usivalidator"
)

func main() {
	prefix := "BNGH7C75F" // Example 9-character prefix
	checkChar, err := usivalidator.GenerateCheckCharacter(prefix)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Printf("The check character is %c\n", checkChar)
	}
}
```

#### IndexOf

```go
package main

import (
	"fmt"

	"github.com/chrisjoyce911/usivalidator"
)

func main() {
	char := 'A'
	index := usivalidator.IndexOf(char, usivalidator.ValidCharacters)
	if index != -1 {
		fmt.Printf("Character %c found at index %d\n", char, index)
	} else {
		fmt.Printf("Character %c not found\n", char)
	}
}
```

#### AlternateFactor

```go
package main

import (
	"fmt"

	"github.com/chrisjoyce911/usivalidator"
)

func main() {
	factor := 2
	nextFactor := usivalidator.AlternateFactor(factor)
	fmt.Printf("Next factor after %d is %d\n", factor, nextFactor)
}
```
