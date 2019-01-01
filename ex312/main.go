package main

import "fmt"

func main() {
	fmt.Println(anagram("hello", "olleh"))
}

func anagram(s1, s2 string) bool {
	r1 := make(map[rune]int)
	r2 := make(map[rune]int)

	for _, r := range s1 {
		r1[r]++
	}

	for _, r := range s2 {
		r2[r]++
	}

	for k, v := range r1 {
		if r2[k] != v {
			return false
		}
	}

	for k, v := range r2 {
		if r1[k] != v {
			return false
		}
	}
	return true
}
