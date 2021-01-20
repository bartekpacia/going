package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	n := len(arr)

	if n < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivotIndex := rand.Int() % len(arr)

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	nums := []int{7, 9, 8, 3, 2, 1, 6, 4, 5}

	fmt.Printf("before: %v\n", nums)
	quickSort(nums)
	fmt.Printf("after : %v\n", nums)
}
