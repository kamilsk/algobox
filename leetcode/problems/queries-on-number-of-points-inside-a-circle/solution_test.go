package main

import (
	"math"
	"reflect"
	"testing"
)

// Solution for https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/.
func TestCountPoints(t *testing.T) {
	tests := []struct {
		points   [][]int
		queries  [][]int
		expected []int
	}{
		{
			points:   [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
			queries:  [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			expected: []int{3, 2, 2},
		},
		{
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			queries:  [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}},
			expected: []int{2, 3, 2, 4},
		},
	}

	for i, test := range tests {
		obtained := countPoints(test.points, test.queries)
		if !reflect.DeepEqual(obtained, test.expected) {
			t.Errorf("expected: %v, obtained: %v (case %d)", test.expected, obtained, i)
		}
	}
}

// n^2
func countPoints(points, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, query := range queries {
		x0, y0, r := query[0], query[1], float64(query[2])

		for _, point := range points {
			x, y := point[0], point[1]

			d := math.Sqrt(float64((x-x0)*(x-x0) + (y-y0)*(y-y0)))
			if r >= d {
				result[i]++
			}
		}
	}

	return result
}
