package main

import "fmt"

func main() {
	d := []int{1, 2, 3}
	fmt.Println(plusOne(d))
}

func plusOne(digits []int) []int {
	l := len(digits)
	newDs := make([]int, l+1)
	newDs[0] = 1
	copy(newDs[1:], digits)
	for i := l; i > 0 && newDs[0] == 1; i-- {
		m := newDs[i] + newDs[0]
		newDs[i] = m % 10
		newDs[0] = m / 10
	}
	if newDs[0] == 0 {
		return newDs[1:]
	}
	return newDs
}
