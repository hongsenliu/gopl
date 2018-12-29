package main

import "fmt"

func main() {
	fmt.Println(getRow(6))
}

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for r := 0; r <= rowIndex; r++ {
		for i := r; i > 0; i-- {
			row[i] = row[i-1] + row[i]
		}
	}
	return row
}
