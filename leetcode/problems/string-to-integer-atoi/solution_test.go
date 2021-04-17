package main

import "testing"

// Solution for https://leetcode.com/problems/string-to-integer-atoi/.
func TestMyAtoi(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},

		{"00000-42a1234", 0},
		{"  0000000000012345678", 12345678},
		{"000000000000000000", 0},
		{"2147483648", 2147483647},
		{"2147483646", 2147483646},
	}

	for i, test := range tests {
		if obtained := myAtoi(test.s); obtained != test.expected {
			t.Errorf("expected: %d, obtained: %d (case %d)",
				test.expected, obtained, i)
		}
	}
}

// n
func myAtoi(s string) int {
	// skip leading whitespaces
	s0 := s
	for _, r := range s0 {
		if r == ' ' {
			s = s[1:]
			continue
		}
		break
	}
	// fail fast
	if len(s) == 0 {
		return 0
	}

	// define sign
	pos := true
	switch s[0] {
	case '-':
		pos = false
		fallthrough
	case '+':
		s = s[1:]
	}

	// skip 0
	s0 = s
	for _, r := range s0 {
		if r == '0' {
			s = s[1:]
			continue
		}
		break
	}

	// clean
	for i, r := range []byte(s) {
		if r-'0' > 9 {
			s = s[:i]
			break
		}
	}
	// fail fast
	if len(s) == 0 {
		return 0
	}

	// now we have numeric sequence
	sLen := len(s)
	if sLen > 10 {
		if pos {
			return 1<<31 - 1
		}
		return -1 << 31
	}
	n := int64(0)
	for _, r := range []byte(s) {
		n = n*10 + int64(r-'0')
	}

	if !pos {
		n = -n
		if n < -1<<31 {
			return -1 << 31
		}
	} else if n > 1<<31-1 {
		return 1<<31 - 1
	}
	return int(n)
}
