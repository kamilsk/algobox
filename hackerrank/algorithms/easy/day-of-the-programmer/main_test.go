package main

import "testing"

func TestDayOfProgrammer(t *testing.T) {
	tests := []struct {
		name     string
		year     int32
		expected string
	}{
		{"Testcase 0", 2017, "13.09.2017"},
		{"Testcase 1", 2016, "12.09.2016"},
		{"Testcase 60", 1800, "12.09.1800"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := dayOfProgrammer(tc.year); tc.expected != obtained {
				t.Errorf("expected %s, obtained %s", tc.expected, obtained)
			}
		})
	}
}
