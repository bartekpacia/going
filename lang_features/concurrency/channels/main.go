package main

import (
	"fmt"
	"time"
)

func main() {
	ichan := make(chan int)

	go func() {
		fmt.Println("goroutine: started")
		// Uncomment 2 lines below and see what happens
		// close(ichan)
		// return

		ichan <- 1
		fmt.Println("goroutine: 1st send received")

		ichan <- 2
		fmt.Println("goroutine: 2nd send received")

		ichan <- 3
		fmt.Println("goroutine: 3rd send received")
	}()

	time.Sleep(1 * time.Second)
	a := <-ichan
	fmt.Println("main: received FIRST value")

	time.Sleep(1 * time.Second)
	b := <-ichan
	fmt.Println("main: received SECOND value")

	time.Sleep(1 * time.Second)
	c := <-ichan
	fmt.Println("main: received THIRD value")

	time.Sleep(1 * time.Second)

	fmt.Println("end! a:", a, "b:", b, "c:", c)
}
