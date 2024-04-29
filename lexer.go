package main

import (
  "fmt"
  "unicode"
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

func tokenize(input string) ([]rune, error) {
  tokens := make([]rune, 0)

  for _, c := range input {
    fmt.Printf("%c", c)
    if !isValidToken(c) && !unicode.IsSpace(c){
      fmt.Printf("\nInvalid token: %c\n", c)
      return nil, fmt.Errorf("Invalid token: %c", c)
    }

    tokens = append(tokens, rune(c))
  }

  return tokens, nil
}
