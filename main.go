package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type AST struct {
	Root   *Node
	Tokens []Token
}

type Node struct {
	Operation   rune
	Value       rune
	Literal     string
	Left, Right *Node
}

func printTree(node *Node, depth int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", depth)

	if node.Literal != "" {
		fmt.Printf("%s%s\n", indent, node.Literal)
	} else if unicode.IsDigit(node.Value) {
		fmt.Printf("%s%c\n", indent, node.Value)
	} else {
		fmt.Printf("%s%c\n", indent, node.Operation)
	}

	printTree(node.Left, depth+2)
	printTree(node.Right, depth+2)
}

func main() {
	fmt.Println("Start tokenization")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	tokenInputs, tokenizeErr := tokenize(input)
	if tokenizeErr != nil {
		fmt.Println("Error tokenizing input", tokenizeErr)
		return
	}

	for _, token := range tokenInputs {
		if token.Type == NUMBER && token.Literal != "" {
			fmt.Printf("Token: Type=%s, Value=%s\n", token.Type, token.Literal)
			continue
		}
		fmt.Printf("Token: Type=%s, Value=%c\n", token.Type, token.Value)
	}

	ast := &AST{
		Tokens: tokenInputs,
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
