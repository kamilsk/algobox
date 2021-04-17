package main

import (
	"strings"
	"testing"
)

// Solution for https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/.
func TestMinPartitions(t *testing.T) {
	tests := []struct {
		n        string
		expected int
	}{
		{"32", 3},
		{"82734", 8},
		{"27346209830709182346", 9},
	}

	for i, test := range tests {
		if obtained := minPartitions(test.n); obtained != test.expected {
			t.Errorf("expected: %d, obtained: %d (case %d)", test.expected, obtained, i)
		}
	}
}

// C*n == n
func minPartitions(n string) int {
	buf := new(strings.Builder)

	var total int
	for n != "" {
		for _, r := range n {
			if r >= '1' {
				buf.WriteRune(r - 1)
				continue
			}
			buf.WriteRune('0')
		}
		m := buf.String()
		n = strings.TrimLeft(m, "0")
		buf.Reset()
		total++
	}
	return total
}
