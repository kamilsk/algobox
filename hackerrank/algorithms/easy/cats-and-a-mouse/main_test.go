package main

import "testing"

func TestCatAndMouse(t *testing.T) {
	tests := []struct {
		name     string
		x, y, z  int32
		expected string
	}{
		{"Testcase 0#1", 1, 2, 3, "Cat B"},
		{"Testcase 0#2", 1, 3, 2, "Mouse C"},
		{"Testcase 1#1", 22, 75, 70, "Cat B"},
		{"Testcase 1#2", 33, 86, 59, "Cat A"},
		{"Testcase 1#3", 47, 29, 89, "Cat A"},
		{"Testcase 1#4", 18, 19, 82, "Cat B"},
		{"Testcase 1#5", 84, 17, 18, "Cat B"},
		{"Testcase 1#6", 100, 11, 55, "Cat B"},
		{"Testcase 1#7", 37, 57, 75, "Cat B"},
		{"Testcase 1#8", 47, 30, 6, "Cat B"},
		{"Testcase 1#9", 40, 68, 50, "Cat A"},
		{"Testcase 1#10", 26, 37, 31, "Cat A"},
		{"Testcase 1#11", 93, 49, 20, "Cat B"},
		{"Testcase 1#12", 84, 78, 31, "Cat B"},
		{"Testcase 1#13", 80, 57, 86, "Cat A"},
		{"Testcase 1#14", 57, 94, 55, "Cat A"},
		{"Testcase 1#15", 21, 80, 4, "Cat A"},
		{"Testcase 1#16", 1, 68, 67, "Cat B"},
		{"Testcase 1#17", 74, 91, 23, "Cat A"},
		{"Testcase 1#18", 85, 66, 80, "Cat A"},
		{"Testcase 1#19", 21, 95, 58, "Mouse C"},
		{"Testcase 1#20", 86, 69, 77, "Cat B"},
		{"Testcase 1#21", 31, 2, 46, "Cat A"},
		{"Testcase 1#22", 45, 94, 99, "Cat B"},
		{"Testcase 1#23", 7, 66, 36, "Cat A"},
		{"Testcase 1#24", 63, 34, 33, "Cat B"},
		{"Testcase 1#25", 75, 92, 65, "Cat A"},
		{"Testcase 1#26", 90, 45, 54, "Cat B"},
		{"Testcase 1#27", 12, 9, 10, "Cat B"},
		{"Testcase 1#28", 43, 56, 51, "Cat B"},
		{"Testcase 1#29", 92, 20, 56, "Mouse C"},
		{"Testcase 1#30", 97, 12, 67, "Cat A"},
		{"Testcase 1#31", 17, 38, 86, "Cat B"},
		{"Testcase 1#32", 85, 94, 20, "Cat A"},
		{"Testcase 1#33", 6, 81, 53, "Cat B"},
		{"Testcase 1#34", 77, 27, 54, "Cat A"},
		{"Testcase 1#35", 62, 25, 37, "Cat B"},
		{"Testcase 1#36", 56, 70, 63, "Mouse C"},
		{"Testcase 1#37", 49, 32, 16, "Cat B"},
		{"Testcase 1#38", 4, 61, 39, "Cat B"},
		{"Testcase 1#39", 92, 30, 61, "Mouse C"},
		{"Testcase 1#40", 41, 59, 81, "Cat B"},
		{"Testcase 1#41", 100, 66, 83, "Mouse C"},
		{"Testcase 1#42", 16, 16, 16, "Mouse C"},
		{"Testcase 1#43", 81, 70, 30, "Cat B"},
		{"Testcase 1#44", 11, 33, 22, "Mouse C"},
		{"Testcase 1#45", 35, 98, 18, "Cat A"},
		{"Testcase 1#46", 43, 62, 48, "Cat A"},
		{"Testcase 1#47", 84, 54, 69, "Mouse C"},
		{"Testcase 1#48", 73, 72, 86, "Cat A"},
		{"Testcase 1#49", 34, 82, 49, "Cat A"},
		{"Testcase 1#50", 16, 83, 62, "Cat B"},
		{"Testcase 1#51", 57, 50, 53, "Cat B"},
		{"Testcase 1#52", 36, 49, 88, "Cat B"},
		{"Testcase 1#53", 5, 80, 42, "Cat A"},
		{"Testcase 1#54", 20, 86, 47, "Cat A"},
		{"Testcase 1#55", 43, 40, 41, "Cat B"},
		{"Testcase 1#56", 72, 12, 42, "Mouse C"},
		{"Testcase 1#57", 16, 43, 29, "Cat A"},
		{"Testcase 1#58", 11, 35, 23, "Mouse C"},
		{"Testcase 1#59", 12, 63, 37, "Cat A"},
		{"Testcase 1#60", 84, 78, 55, "Cat B"},
		{"Testcase 1#61", 17, 90, 78, "Cat B"},
		{"Testcase 1#62", 28, 10, 84, "Cat A"},
		{"Testcase 1#63", 39, 96, 67, "Cat A"},
		{"Testcase 1#64", 22, 84, 53, "Mouse C"},
		{"Testcase 1#65", 49, 77, 63, "Mouse C"},
		{"Testcase 1#66", 77, 82, 55, "Cat A"},
		{"Testcase 1#67", 17, 53, 35, "Mouse C"},
		{"Testcase 1#68", 79, 31, 55, "Mouse C"},
		{"Testcase 1#69", 7, 56, 31, "Cat A"},
		{"Testcase 1#70", 2, 7, 4, "Cat A"},
		{"Testcase 1#71", 99, 82, 60, "Cat B"},
		{"Testcase 1#72", 20, 17, 18, "Cat B"},
		{"Testcase 1#73", 1, 98, 49, "Cat A"},
		{"Testcase 1#74", 91, 66, 13, "Cat B"},
		{"Testcase 1#75", 95, 23, 1, "Cat B"},
		{"Testcase 1#76", 87, 59, 73, "Mouse C"},
		{"Testcase 1#77", 10, 10, 56, "Mouse C"},
		{"Testcase 1#78", 61, 54, 59, "Cat A"},
		{"Testcase 1#79", 62, 94, 78, "Mouse C"},
		{"Testcase 1#80", 49, 29, 37, "Cat B"},
		{"Testcase 1#81", 87, 79, 83, "Mouse C"},
		{"Testcase 1#82", 72, 1, 42, "Cat A"},
		{"Testcase 1#83", 42, 34, 38, "Mouse C"},
		{"Testcase 1#84", 52, 82, 98, "Cat B"},
		{"Testcase 1#85", 29, 12, 43, "Cat A"},
		{"Testcase 1#86", 81, 50, 97, "Cat A"},
		{"Testcase 1#87", 80, 17, 43, "Cat B"},
		{"Testcase 1#88", 88, 38, 40, "Cat B"},
		{"Testcase 1#89", 41, 55, 84, "Cat B"},
		{"Testcase 1#90", 48, 91, 69, "Cat A"},
		{"Testcase 1#91", 11, 74, 23, "Cat A"},
		{"Testcase 1#92", 84, 68, 76, "Mouse C"},
		{"Testcase 1#93", 4, 51, 80, "Cat B"},
		{"Testcase 1#94", 51, 85, 39, "Cat A"},
		{"Testcase 1#95", 37, 74, 55, "Cat A"},
		{"Testcase 1#96", 15, 65, 54, "Cat B"},
		{"Testcase 1#97", 57, 14, 56, "Cat A"},
		{"Testcase 1#98", 43, 61, 56, "Cat B"},
		{"Testcase 1#99", 9, 79, 35, "Cat A"},
		{"Testcase 1#100", 4, 44, 44, "Cat B"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := catAndMouse(tc.x, tc.y, tc.z); tc.expected != obtained {
				t.Errorf("expected %s, obtained %s", tc.expected, obtained)
			}
		})
	}
}
