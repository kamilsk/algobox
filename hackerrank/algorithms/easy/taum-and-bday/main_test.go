package main

import "testing"

func TestTaumBday(t *testing.T) {
	tests := []struct {
		name            string
		b, w, bc, wc, z int64
		expected        int64
	}{
		{"Test case 0#1", 10, 10, 1, 1, 1, 20},
		{"Test case 0#2", 5, 9, 2, 3, 4, 37},
		{"Test case 0#3", 3, 6, 9, 1, 1, 12},
		{"Test case 0#4", 7, 7, 4, 2, 1, 35},
		{"Test case 0#5", 3, 3, 1, 9, 2, 12},
		{"Test case 5#1", 27984, 1402, 619246, 615589, 247954, 18192035842},
		{"Test case 5#2", 95677, 39394, 86983, 311224, 588538, 20582630747},
		{"Test case 5#3", 68406, 12718, 532909, 315341, 201009, 39331944938},
		{"Test case 5#4", 15242, 95521, 712721, 628729, 396706, 70920116291},
		{"Test case 5#5", 21988, 67781, 645748, 353796, 333199, 38179353700},
		{"Test case 5#6", 22952, 80129, 502975, 175136, 340236, 25577754744},
		{"Test case 5#7", 38577, 3120, 541637, 657823, 735060, 22947138309},
		{"Test case 5#8", 5943, 69851, 674453, 392925, 381074, 31454478354},
		{"Test case 5#9", 62990, 61330, 310144, 312251, 93023, 38686324390},
		{"Test case 5#10", 11152, 43844, 788543, 223872, 972572, 18609275504},
		{"Test case 14", 3, 5, 3, 4, 1, 29},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := taumBday(tc.b, tc.w, tc.bc, tc.wc, tc.z); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
