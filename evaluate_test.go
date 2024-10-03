package main

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	testCases := []struct {
		name     string
		input    []Token
		expected int
	}{
		{
			name: "addition",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '1'},
				{Type: NUMBER, Value: '2'},
				{Type: RPAREN, Value: ')'},
			},
			expected: 3,
		},
		{
			name: "multiplication",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: MULT, Value: '*'},
				{Type: NUMBER, Value: '3'},
				{Type: NUMBER, Value: '4'},
				{Type: RPAREN, Value: ')'},
			},
			expected: 12,
		},
		{
			name: "subtraction",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: MINUS, Value: '-'},
				{Type: NUMBER, Value: '5'},
				{Type: NUMBER, Value: '6'},
				{Type: RPAREN, Value: ')'},
			},
			expected: -1,
		},
		{
			name: "division",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: DIV, Value: '/'},
				{Type: NUMBER, Value: '8'},
				{Type: NUMBER, Value: '2'},
				{Type: RPAREN, Value: ')'},
			},
			expected: 4,
		},
		{
			name: "nested expression",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '1'},
				{Type: LPAREN, Value: '('},
				{Type: MULT, Value: '*'},
				{Type: NUMBER, Value: '2'},
				{Type: NUMBER, Value: '3'},
				{Type: RPAREN, Value: ')'},
				{Type: RPAREN, Value: ')'},
			},
			expected: 7,
		},
		{
			name: "complex nested expression",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: MINUS, Value: '-'},
				{Type: LPAREN, Value: '('},
				{Type: MULT, Value: '*'},
				{Type: NUMBER, Value: '3'},
				{Type: NUMBER, Value: '4'},
				{Type: RPAREN, Value: ')'},
				{Type: LPAREN, Value: '('},
				{Type: DIV, Value: '/'},
				{Type: NUMBER, Value: '8'},
				{Type: NUMBER, Value: '2'},
				{Type: RPAREN, Value: ')'},
				{Type: RPAREN, Value: ')'},
			},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ast := &AST{
				Tokens: tc.input,
			}

			_, err := ast.Parse(0, ast)
			if err != nil {
				t.Fatalf("Parse error for input %v: %v", tc.input, err)
			}

			result := ast.Evaluate(nil)
			if result != tc.expected {
				t.Errorf("For input %v, expected %d but got %d", tc.input, tc.expected, result)
			}
		})
	}
}
