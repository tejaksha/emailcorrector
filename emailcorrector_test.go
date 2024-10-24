package emailcorrector

import (
	"fmt"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"user.name+tag@gmail.com",
	}

	for _, email := range validEmails {
		if err := ValidateEmail(email); err != nil {
			t.Errorf("expected valid email, got error: %v", err)
		}
	}

	invalidEmails := []string{
		"invalidemail.com",
		"user@",
		"@example.com",
	}

	for _, email := range invalidEmails {
		if err := ValidateEmail(email); err == nil {
			t.Errorf("expected error for invalid email, got none")
		}
	}
}

func TestSuggestDomainCorrection(t *testing.T) {
	tests := map[string]string{
		"gmial.com":  "gmail.com",
		"yaho.com":   "yahoo.com",
		"outlok.com": "outlook.com",
	}

	for input, expected := range tests {
		result := SuggestDomainCorrection(input)
		fmt.Println(result)
		if result != expected {
			t.Errorf("expected %s, got %s", expected, result)
		}
	}
}

func TestCorrectEmail(t *testing.T) {
	email := "user@gmial.com"
	corrected, err := CorrectEmail(email)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	fmt.Println(corrected)
	expected := "user@gmail.com"
	if corrected != expected {
		t.Errorf("expected %s but got %s", expected, corrected)
	}
}
