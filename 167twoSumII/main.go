package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 25}, 9))
}

func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	for p := 0; p < l; p++ {
		t := target - numbers[p]
		if q := findTarget(numbers[p+1:], t); q != -1 {
			return []int{p + 1, q + 2}
		}
	}
	return nil
}

func findTarget(numbers []int, target int) int {
	l, r := 0, len(numbers)-1
	for l <= r {
		m := l + (r-l)/2
		if numbers[m] == target {
			return m
		}
		if numbers[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
