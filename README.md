# Emailcorrector

A Go package for validating and auto-correcting email addresses.

## Features

- Validate the format of an email address.
- Suggest corrections for common domain name typos (e.g., `gmial.com` → `gmail.com`).
- Use the Levenshtein distance algorithm to detect and correct minor mistakes.

## Installation

```bash
go get github.com/tejaksha/emailcorrector
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/tejaksha/emailcorrector"
)

func main() {
    // Example of validating an email
    email := "test@example.com"
    if err := emailcorrector.ValidateEmail(email); err != nil {
        fmt.Println("Invalid email:", err)
    } else {
        fmt.Println("Valid email:", email)
    }

    // Example of suggesting a domain correction
    incorrectDomain := "gmial.com"
    suggested := emailcorrector.SuggestDomainCorrection(incorrectDomain)
    fmt.Printf("Did you mean: %s?
", suggested)

    // Example of correcting an email
    incorrectEmail := "user@gmial.com"
    corrected, err := emailcorrector.CorrectEmail(incorrectEmail)
    if err != nil {
        fmt.Println("Invalid email:", err)
    } else {
        fmt.Println("Corrected email:", corrected)
    }
}
```

## License

MIT License
