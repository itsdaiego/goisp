package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestParse(t *testing.T) {
	case1 := []Token{
		{Type: LPAREN, Value: '('},
		{Type: PLUS, Value: '+'},
		{Type: NUMBER, Value: '1'},
		{Type: NUMBER, Value: '2'},
		{Type: RPAREN, Value: ')'},
	}
	case2 := []Token{
		{Type: LPAREN, Value: '('},
		{Type: MULT, Value: '*'},
		{Type: NUMBER, Value: '3'},
		{Type: NUMBER, Value: '4'},
		{Type: RPAREN, Value: ')'},
	}
	case3 := []Token{
		{Type: LPAREN, Value: '('},
		{Type: MINUS, Value: '-'},
		{Type: NUMBER, Value: '5'},
		{Type: NUMBER, Value: '6'},
		{Type: RPAREN, Value: ')'},
	}
	case4 := []Token{
    {Type: LPAREN, Value: '('},
    {Type: DIV, Value: '/'},
		{Type: NUMBER, Value: '8'},
		{Type: NUMBER, Value: '2'},
		{Type: RPAREN, Value: ')'},
	}
	case5 := []Token{
		{Type: LPAREN, Value: '('},
		{Type: PLUS, Value: '+'},
		{Type: LPAREN, Value: '('},
		{Type: MULT, Value: '*'},
		{Type: NUMBER, Value: '2'},
		{Type: NUMBER, Value: '3'},
		{Type: RPAREN, Value: ')'},
		{Type: LPAREN, Value: '('},
		{Type: MINUS, Value: '-'},
		{Type: NUMBER, Value: '5'},
		{Type: NUMBER, Value: '1'},
		{Type: RPAREN, Value: ')'},
		{Type: RPAREN, Value: ')'},
	}
	testCases := []struct {
		input    []Token
		expected string
	}{
		{case1, "+\n  1\n  2\n"},
		{case2, "*\n  3\n  4\n"},
		{case3, "-\n  5\n  6\n"},
		{case4, "/\n  8\n  2\n"},
		{case5, "+\n  *\n    2\n    3\n  -\n    5\n    1\n"},
	}

	for _, tc := range testCases {
		ast := &AST{
			Tokens: tc.input,
		}

		_, err := ast.Parse(0, ast)
		if err != nil {
			t.Errorf("Parse error for input %v: %v", tc.input, err)
			continue
		}

		var output strings.Builder
		printTreeToBuilder(ast.Root, 0, &output)
		if output.String() != tc.expected {
			t.Errorf("For input %v, expected:\n%s\nbut got:\n%s", tc.input, tc.expected, output.String())
		}
	}
}

func printTreeToBuilder(node *Node, depth int, builder *strings.Builder) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", depth)
	if !unicode.IsDigit(node.Value) {
		builder.WriteString(fmt.Sprintf("%s%c\n", indent, node.Value))
	} else {
		builder.WriteString(fmt.Sprintf("%s%c\n", indent, node.Operation))
	}

	printTreeToBuilder(node.Left, depth+1, builder)
	printTreeToBuilder(node.Right, depth+1, builder)
}
