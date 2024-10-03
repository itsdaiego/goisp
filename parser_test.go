package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name     string
		input    []Token
		expected *Node
		hasError bool
	}{
		{
			name: "Simple addition",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '1'},
				{Type: NUMBER, Value: '2'},
				{Type: RPAREN, Value: ')'},
			},
			expected: &Node{
				Operation: '+',
				Left:      &Node{Value: '1'},
				Right:     &Node{Value: '2'},
			},
			hasError: false,
		},
		{
			name: "Nested expression",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: MULT, Value: '*'},
				{Type: NUMBER, Value: '3'},
				{Type: LPAREN, Value: '('},
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '4'},
				{Type: NUMBER, Value: '5'},
				{Type: RPAREN, Value: ')'},
				{Type: RPAREN, Value: ')'},
			},
			expected: &Node{
				Operation: '*',
				Left:      &Node{Value: '3'},
				Right: &Node{
					Operation: '+',
					Left:      &Node{Value: '4'},
					Right:     &Node{Value: '5'},
				},
			},
			hasError: false,
		},
		{
			name: "Missing opening parenthesis",
			input: []Token{
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '1'},
				{Type: NUMBER, Value: '2'},
				{Type: RPAREN, Value: ')'},
			},
			expected: nil,
			hasError: true,
		},
		{
			name: "Missing closing parenthesis",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '1'},
				{Type: NUMBER, Value: '2'},
			},
			expected: nil,
			hasError: true,
		},
		{
			name: "Invalid operator",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: NUMBER, Value: '1'},
				{Type: NUMBER, Value: '2'},
				{Type: NUMBER, Value: '3'},
				{Type: RPAREN, Value: ')'},
			},
			expected: nil,
			hasError: true,
		},
		{
			name: "Too many operands",
			input: []Token{
				{Type: LPAREN, Value: '('},
				{Type: PLUS, Value: '+'},
				{Type: NUMBER, Value: '1'},
				{Type: NUMBER, Value: '2'},
				{Type: NUMBER, Value: '3'},
				{Type: RPAREN, Value: ')'},
			},
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ast := &AST{Tokens: tc.input}
			_, err := ast.Parse(0, ast)

			if tc.hasError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if !reflect.DeepEqual(ast.Root, tc.expected) {
					t.Errorf("Expected AST:\n%v\nGot:\n%v", tc.expected, ast.Root)
				}
			}
		})
	}
}

func compareNodes(n1, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	return n1.Operation == n2.Operation &&
		n1.Value == n2.Value &&
		compareNodes(n1.Left, n2.Left) &&
		compareNodes(n1.Right, n2.Right)
}
