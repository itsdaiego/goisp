package main

import (
	"fmt"
	"strconv"
)

func evaluate(node *Node) int {
	if node == nil {
		return 0
	}

  fmt.Println("Source value", string(node.Operation))

	if node.Left == nil && node.Right == nil {
		if num, err := strconv.Atoi(string(node.Value)); err == nil {
			return num
		}
		return 0
	}

	leftVal := evaluate(node.Left)
	rightVal := evaluate(node.Right)

	switch node.Operation {
	case '+':
		return leftVal + rightVal
	case '-':
		return leftVal - rightVal
	case '*':
		return leftVal * rightVal
	case '/':
		if rightVal == 0 {
			fmt.Println("That doesnt make senseee buddy!")
			return 0
		}
		return leftVal / rightVal
	default:
		fmt.Printf("Unsupported operation '%c'\n", node.Operation)
		return 0
	}
}
