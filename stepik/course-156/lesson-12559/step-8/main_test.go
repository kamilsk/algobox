package main

import "testing"

func TestFix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"case#01", "}[[([{[]}", "{}[[([{[]}])]]"},
		{"case#02", "{[({[({[(", "{[({[({[()]})]})]}"},
		{"case#03", ")]})]})]}", "{[({[({[()]})]})]}"},
		{"case#04", "{][[[[{}[]", "IMPOSSIBLE"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := fix(tc.input); tc.expected != obtained {
				t.Errorf("expected %s, obtained %s", tc.expected, obtained)
			}
		})
	}
}
