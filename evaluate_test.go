package main

import (
  "testing"
)

func TestEvaluate(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"(+ 1 2)", 3},
		{"(* 3 4)", 12},
		{"(- 5 6)", -1},
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

		result := ast.Evaluate(nil)
		if result != tc.expected {
			t.Errorf("For input %s, expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}
