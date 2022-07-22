package challenge06

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name     string
	input    string
	expected bool
}

var testCases []TestCase = []TestCase{
	{
		name:     "Valid parentheses order",
		input:    "(()())",
		expected: true,
	},
	{
		name:     "Invalid parentheses order, different number of open and closed parentheses",
		input:    "()()())",
		expected: false,
	},
	{
		name:     "Invalid parentheses order, same number of open and closed parentheses, but in wrong order",
		input:    "()()())(",
		expected: false,
	},
	{
		name:     "Valid, but big parentheses order",
		input:    "((()))(())",
		expected: true,
	},
}

func TestValidateParenthesesOrder(t *testing.T) {
	for _, tc := range testCases {
		output := ValidateParenthesesOrder(tc.input)

		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "int")
		}
	}
}

func TestValidateParenthesesOrderOpt(t *testing.T) {
	for _, tc := range testCases {
		output := ValidateParenthesesOrderOpt(tc.input)

		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("%s got %v, wanted %v, of type %s", tc.name, output, tc.expected, "int")
		}
	}
}
