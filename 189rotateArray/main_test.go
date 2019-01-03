package main

import "testing"

func TestRotate(t *testing.T) {
	var tests = []struct {
		input []int
		k     int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, test := range tests {
		rotate(test.input, test.k)
		for i, v := range test.input {
			if v != test.want[i] {
				t.Errorf("actual: %d, expected: %d", v, test.want[i])
			}
		}
	}
}
