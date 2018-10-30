package main

import "testing"

func TestCheck(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"case#01", "[]", -1},
		{"case#02", "{}[]", -1},
		{"case#03", "[()]", -1},
		{"case#04", "(())", -1},
		{"case#05", "{[]}()", -1},
		{"case#06", "{", 1},
		{"case#07", "{[}", 3},
		{"case#08", "foo(bar);", -1},
		{"case#09", "foo(bar[i);", 10},
		{"case#10", "([](){([])})", -1},
		{"case#11", "()[]}", 5},
		{"case#12", "{{[()]]", 7},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := check(tc.input); obtained != tc.expected {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
