package main

import (
	"fmt"
	"strings"
)

type AST struct {
	Root *Node
}

type Node struct {
	Operation   rune
  Value rune
	Left, Right *Node
}

func printTree(node *Node, depth int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", depth)
  if (isOperation(node.Operation)) {
    fmt.Printf("%s%c\n", indent, node.Operation)
  } else {
    fmt.Printf("%s%c\n", indent, node.Value)
  }

	printTree(node.Left, depth+2)
	printTree(node.Right, depth+2)
}

func main() {
	fmt.Println("Start tokenization")

	// tokenRunes, tokenError := tokenize("(- (+ (+ 1 4) 5) 5)")
	// tokenRunes, tokenError := tokenize("(+ 2 (- 4 (+ 3 4)) )")
	tokenRunes, tokenError := tokenize("(+ 2 (* 5 5))")

	if tokenError != nil {
		fmt.Println("Error tokenizing input", tokenError)

		return
	}

	fmt.Println("\nStart parsing")

	ast := &AST{}

	_, parseError := parse(tokenRunes, 0, ast)

	if parseError != nil {
		fmt.Println("Error parsing tokens", parseError)

		return
	}

	printTree(ast.Root, 0)

	fmt.Println("Syntax is correct!")

  fmt.Println("\nStart evaluating")

  result := evaluate(ast.Root)

  fmt.Println("Result:", result)
}
