package main

import "fmt"

func main() {
	arr := []int{}

	for {
		var v int
		_, err := fmt.Scanf("%d", &v)
		if err != nil {
			break
		}

		arr = append(arr, v)
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println("sum:", sum)
}
