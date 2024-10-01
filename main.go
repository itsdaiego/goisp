package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AST struct {
	Root *Node
  Tokens []rune
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

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	tokenRunes := []rune(input)

  ast := &AST{
    Tokens: tokenRunes,
  }

	_, parseError := ast.Parse(0, ast)

	if parseError != nil {
		fmt.Println("Error parsing tokens", parseError)

		return
	}

	printTree(ast.Root, 0)

	fmt.Println("Syntax is correct!")

	fmt.Println("\nStart evaluating")

	result := ast.Evaluate(nil)

	fmt.Println("Result:", result)
}
