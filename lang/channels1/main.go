package main

import (
	"fmt"
	"time"
)

func main() {
	publish("Jebać pis!", 5*time.Second)
	fmt.Println("Ciekawe co tam słychać w Polsce...")

	time.Sleep(10 * time.Second)
	fmt.Println("Prawidłowo!")
}

func publish(news string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(news)
	}()
}
