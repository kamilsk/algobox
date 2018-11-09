package main

import "testing"

func TestTimeConversion(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"Testcase 0", "07:05:45PM", "19:05:45"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := timeConversion(tc.s); tc.expected != obtained {
				t.Errorf("expected %s, obtained %s", tc.expected, obtained)
			}
		})
	}
}
