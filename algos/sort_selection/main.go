package main

import "fmt"

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}

func main() {
	nums := []int{7, 9, 8, 3, 2, 1, 6, 4, 5}

	fmt.Printf("before: %v\n", nums)
	selectionSort(nums)
	fmt.Printf("after : %v\n", nums)
}
