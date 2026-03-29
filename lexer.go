package main

import (
	"fmt"
	"unicode"
)

type Token struct {
	Type    string
	Value   rune
	Literal string
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

func isValidNumber(token string) bool {
	runes := []rune(token)

	for _, r := range runes {
		if !unicode.IsNumber(r) || unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func tokenize(input string) ([]Token, error) {
	tokens := make([]Token, 0)

	runes := []rune(input)
	for i := 0; i < len(runes); {
		if unicode.IsSpace(runes[i]) {
			i++
			continue
		}

		current := runes[i]

		switch current {
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: '('})
			i++
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ')'})
			i++
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Value: '+'})
			i++
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Value: '-'})
			i++
		case '*':
			tokens = append(tokens, Token{Type: MULT, Value: '*'})
			i++
		case '/':
			tokens = append(tokens, Token{Type: DIV, Value: '/'})
			i++
		default:
			if !isValidToken(runes[i]) {
				return nil, fmt.Errorf("invalid token: %c", runes[i])
			}

			start := i
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				i++
			}

			number := string(runes[start:i])
			if !isValidNumber(number) {
				return nil, fmt.Errorf("invalid token: %s", number)
			}

			token := Token{Type: NUMBER, Value: runes[start]}
			if len(runes[start:i]) > 1 {
				token.Literal = number
			}

			tokens = append(tokens, token)
		}
	}

	return tokens, nil
}
