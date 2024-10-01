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

func (ast *AST) Parse(index int, parent interface{}) (int, error) {
	// every expression should start with a '('
	if ast.Tokens[index] != ParenOpen {
		return 0, fmt.Errorf("Syntax Error: Expected '(', found '%c'", ast.Tokens[index])
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

	// we increase to what we expect to be a operation
	index++

	if index == len(ast.Tokens) || !isOperation(ast.Tokens[index]) {
		return 0, fmt.Errorf("Syntax Error: Expected operator after '(', found '%c'", ast.Tokens[index])
	}

	node.Operation = ast.Tokens[index]

	// we increase it again to start parsing the operands
	// or the next '('
	index++
	operandCount := 0

	for index < len(ast.Tokens) && ast.Tokens[index] != ParenClose {
		if ast.Tokens[index] == ParenOpen {
			var err error
			index, err = ast.Parse(index, node)
			if err != nil {
				return 0, err
			}
		} else if unicode.IsDigit(ast.Tokens[index]) {
			digitNode := &Node{Value: ast.Tokens[index]}
			if operandCount == 0 {
				node.Left = digitNode
			} else {
				node.Right = digitNode
			}
			index++
		} else if unicode.IsSpace(ast.Tokens[index]) {
			index++
			continue
		} else {
			return 0, fmt.Errorf("Syntax Error: Invalid token '%c' in expression", ast.Tokens[index])
		}

		operandCount++
	}

	if index == len(ast.Tokens) || ast.Tokens[index] != ParenClose {
		return 0, fmt.Errorf("Syntax Error: Expected ')'")
	}

	if operandCount != 2 {
		return 0, fmt.Errorf("Syntax Error: Operator requires exactly two operands, found %d", operandCount)
	}

	return index + 1, nil
}
