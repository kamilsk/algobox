package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			t.Log(tc.name)
		})
	}
}
