package main

import "fmt"

// Source: https://austingwalters.com/merge-sort-in-go-golang/#disclosure

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	arr := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			arr[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			arr[k] = left[i]
			i++
		} else if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	arr1 := mergeSort(arr[:mid])
	arr2 := mergeSort(arr[mid:])

	return merge(arr1, arr2)
}

func main() {
	nums := []int{7, 9, 8, 3, 2, 1, 6, 4, 5}

	fmt.Printf("before: %v\n", nums)
	nums = mergeSort(nums)
	fmt.Printf("after : %v\n", nums)
}
