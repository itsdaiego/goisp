package main

import (
	"fmt"
	"unicode"
)

var (
	ParenOpen  = '('
	ParenClose = ')'
	operations = []rune{'+', '-', '*', '/'}
)

func isOperation(token rune) bool {
	for _, op := range operations {
		if token == op {
			return true
		}
	}
	return false
}

func parse(tokens []rune, index int, parent interface{}) (int, error) {
	if tokens[index] != ParenOpen {
		return 0, fmt.Errorf("Syntax Error: Expected '(', found '%c'", tokens[index])
	}

	node := &Node{}

	if ast, ok := parent.(*AST); ok {
		ast.Root = node
	} else if parentNode, ok := parent.(*Node); ok {
		if parentNode.Left == nil {
			parentNode.Left = node
		} else {
			parentNode.Right = node
		}
	}

	index++

	if index == len(tokens) || !isOperation(tokens[index]) {
		return 0, fmt.Errorf("Syntax Error: Expected operator after '(', found '%c'", tokens[index])
	}

  fmt.Println("PARSING", string(tokens[index]))
	node.Operation = tokens[index]

	index++
	operandCount := 0

	for index < len(tokens) && tokens[index] != ParenClose {
		if tokens[index] == ParenOpen {
			var err error
			index, err = parse(tokens, index, node)
			if err != nil {
				return 0, err
			}
		} else if unicode.IsDigit(tokens[index]) {
			digitNode := &Node{Value: tokens[index]}
			if operandCount == 0 {
				node.Left = digitNode
			} else {
				node.Right = digitNode
			}
			index++
		} else {
			return 0, fmt.Errorf("Syntax Error: Invalid token '%c' in expression", tokens[index])
		}

		operandCount++
	}

	if index == len(tokens) || tokens[index] != ParenClose {
		return 0, fmt.Errorf("Syntax Error: Expected ')'")
	}

	if operandCount != 2 {
		return 0, fmt.Errorf("Syntax Error: Operator requires exactly two operands, found %d", operandCount)
	}

	return index + 1, nil
}
