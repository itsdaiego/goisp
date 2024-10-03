package main

import (
	"fmt"
	"unicode"
)


type Token struct {
	Type  string
	Value rune
}


const (
	LPAREN string = "LPAREN"
	RPAREN string = "RPAREN"
	PLUS   string = "PLUS"
	MINUS  string = "MINUS"
	MULT   string = "MULT"
	DIV    string = "DIV"
	NUMBER string = "NUMBER"
)

func isValidToken(token rune) bool {
	if token == ')' || token == '(' {
		return true
	}

	operations := []rune{'+', '-', '*', '/'}

	for _, operation := range operations {
		if token == operation {
			return true
		}
	}

	return unicode.IsDigit(token)
}

func tokenize(input string) ([]Token, error) {
	tokens := make([]Token, 0)

	for i, c := range input {
		if unicode.IsSpace(c) {
			continue
		}

		if !isValidToken(c) {
			return nil, fmt.Errorf("Invalid token: %c", c)
		}

		switch c {
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: '('})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ')'})
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Value: '+'})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Value: '-'})
		case '*':
			tokens = append(tokens, Token{Type: MULT, Value: '*'})
		case '/':
			tokens = append(tokens, Token{Type: DIV, Value: '/'})
		default:
			if unicode.IsDigit(c) {
				tokens = append(tokens, Token{Type: NUMBER, Value: c})
			}
		}
	}

	return tokens, nil
}
