package main

import "fmt"

func main() {
	nums, target := []int{1}, 0
	i := searchInsert(nums, target)
	fmt.Println(i)
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	m := l + (r-l)/2
	for l <= r {
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
		m = l + (r-l)/2
	}
	return m
}
