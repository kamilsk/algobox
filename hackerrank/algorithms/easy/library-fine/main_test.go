package main

import "testing"

func TestLibraryFine(t *testing.T) {
	tests := []struct {
		name       string
		d1, m1, y1 int32
		d2, m2, y2 int32
		expected   int32
	}{
		{"Test case 1", 9, 6, 2015, 6, 6, 2015, 45},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := libraryFine(tc.d1, tc.m1, tc.y1, tc.d2, tc.m2, tc.y2); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
