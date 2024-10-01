package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"(+ 1 2)", "+\n  1\n  2\n"},
		{"(* 3 4)", "*\n  3\n  4\n"},
		{"(- 5 6)", "-\n  5\n  6\n"},
		{"(/ 8 2)", "/\n  8\n  2\n"},
		{"(+ (* 2 3) (- 5 1))", "+\n  *\n    2\n    3\n  -\n    5\n    1\n"},
	}

	for _, tc := range testCases {
		ast := &AST{
			Tokens: []rune(tc.input),
		}

		_, err := ast.Parse(0, ast)
		if err != nil {
			t.Errorf("Parse error for input %s: %v", tc.input, err)
			continue
		}

		var output strings.Builder
		printTreeToBuilder(ast.Root, 0, &output)
		if output.String() != tc.expected {
			t.Errorf("For input %s, expected:\n%s\nbut got:\n%s", tc.input, tc.expected, output.String())
		}
	}
}

func printTreeToBuilder(node *Node, depth int, builder *strings.Builder) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", depth)
	if isOperation(node.Operation) {
		builder.WriteString(fmt.Sprintf("%s%c\n", indent, node.Operation))
	} else {
		builder.WriteString(fmt.Sprintf("%s%c\n", indent, node.Value))
	}

	printTreeToBuilder(node.Left, depth+1, builder)
	printTreeToBuilder(node.Right, depth+1, builder)
}
