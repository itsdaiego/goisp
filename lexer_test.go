package main

import (
	"reflect"
	"testing"
)

var TOKENS_MAP = map[rune]string{
	'(': "ParenOpen",
	')': "ParenClose",
	'+': "Plus",
}

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

func TestTokenize(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Token
		hasError bool
	}{
		{
			input: "(+ 1 2)",
			expected: []Token{
				{Type: LPAREN, Value: '('},
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '1'},
				{Type: NUMBER, Value: '2'},
				{Type: RPAREN, Value: ')'},
			},
			hasError: false,
		},
		{
			input: "(* 3 4)",
			expected: []Token{
				{Type: LPAREN, Value: '('},
				{Type: MULT, Value: '*'},
				{Type: NUMBER, Value: '3'},
				{Type: NUMBER, Value: '4'},
				{Type: RPAREN, Value: ')'},
			},
			hasError: false,
		},
		{
			input:    "invalid input @",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		result, err := tokenize(tc.input)

		if tc.hasError {
			if err == nil {
				t.Errorf("Expected error for input %s, but got none", tc.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tc.input, err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For input %s, expected %v but got %v", tc.input, tc.expected, result)
			}
		}
	}
}
