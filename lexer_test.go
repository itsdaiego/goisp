package main

import (
	"testing"
)

func TestIsValidToken(t *testing.T) {
	validTokens := []rune{'(', ')', '+', '-', '*', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	invalidTokens := []rune{'a', 'b', 'c', '@', '#', '$', '%', '^', '&', '_'}

	for _, token := range validTokens {
		if !isValidToken(token) {
			t.Errorf("Token %c should be valid", token)
		}
	}

	for _, token := range invalidTokens {
		if isValidToken(token) {
			t.Errorf("Token %c should be invalid", token)
		}
	}
}
