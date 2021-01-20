package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	nums := []int{7, 9, 8, 3, 2, 1, 6, 4, 5}

	fmt.Printf("before: %v\n", nums)
	bubbleSort(nums)
	fmt.Printf("after : %v\n", nums)
}
