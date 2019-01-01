package main

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
	}

	for _, test := range tests {
		if got := twoSum(test.input, test.target); got[0] != test.want[0] || got[1] != test.want[1] {
			t.Errorf("actual: %v, want: %v", got, test.want)
		}
	}
}
