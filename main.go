package main

import (
  "fmt"
)

func main() {
  fmt.Println("Start tokenization")

  tokenRunes, error := tokenize("(+ 1 2)")

  if (error != nil) {
    fmt.Println("Error tokenizing input", error)

    return
  }

  tokens := string(tokenRunes)

  fmt.Println("\nshowing tokens", tokens)
}
