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

func (ast *AST) Parse(index int, parent interface{}) (int, error) {
	// every expression should start with a '('
	if ast.Tokens[index].Type != LPAREN {
		return 0, fmt.Errorf("Syntax Error: Expected '(', found '%c'", ast.Tokens[index].Value)
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
		return 0, fmt.Errorf("Syntax Error: Expected operator after '(', found '%c'", ast.Tokens[index].Value)
	}

  if (isOperation(ast.Tokens[index])) {
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
			digitNode := &Node{Value: ast.Tokens[index].Value}
			if operandCount == 0 {
				node.Left = digitNode
			} else {
				node.Right = digitNode
			}
			index++
		} else {
			return 0, fmt.Errorf("Syntax Error: Invalid token '%c' in expression", ast.Tokens[index].Value)
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
