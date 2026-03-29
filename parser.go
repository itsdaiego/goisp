package main

import (
	"fmt"
)

func isOperation(token Token) bool {
	var operations = []rune{'+', '-', '*', '/'}
	for _, op := range operations {
		if token.Value == op {
			return true
		}
	}
	return false
}

func tokenText(token Token) string {
	if token.Type == NUMBER && token.Literal != "" {
		return token.Literal
	}
	if token.Value == 0 {
		return ""
	}
	return string(token.Value)
}

func (ast *AST) Parse(index int, parent interface{}) (int, error) {
	if len(ast.Tokens) == 0 || index >= len(ast.Tokens) {
		return 0, fmt.Errorf("syntax error: empty or incomplete expression")
	}

	// every expression should start with a '('
	if ast.Tokens[index].Type != LPAREN {
		return 0, fmt.Errorf("Syntax Error: Expected '(', found '%s'", tokenText(ast.Tokens[index]))
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
		found := "EOF"
		if index < len(ast.Tokens) {
			found = tokenText(ast.Tokens[index])
		}
		return 0, fmt.Errorf("Syntax Error: Expected operator after '(', found '%s'", found)
	}

	if isOperation(ast.Tokens[index]) {
		node.Operation = ast.Tokens[index].Value
	}

	// we increase it again to start parsing the operands
	// or the next '('
	index++
	operandCount := 0

	for index < len(ast.Tokens) && ast.Tokens[index].Type != RPAREN {
		if ast.Tokens[index].Type == LPAREN {
			var err error
			index, err = ast.Parse(index, node)
			if err != nil {
				return 0, err
			}
		} else if ast.Tokens[index].Type == NUMBER {
			literal := ast.Tokens[index].Literal
			digitNode := &Node{Value: ast.Tokens[index].Value, Literal: literal}
			if operandCount == 0 {
				node.Left = digitNode
			} else {
				node.Right = digitNode
			}
			index++
		} else {
			return 0, fmt.Errorf("Syntax Error: Invalid token '%s' in expression", tokenText(ast.Tokens[index]))
		}

		operandCount++
	}

	if index == len(ast.Tokens) || ast.Tokens[index].Type != RPAREN {
		return 0, fmt.Errorf("Syntax Error: Expected ')'")
	}

	if operandCount != 2 {
		return 0, fmt.Errorf("Syntax Error: Operator requires exactly two operands, found %d", operandCount)
	}

	return index + 1, nil
}
