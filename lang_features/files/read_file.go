package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}

	buffer := make([]byte, 33)
	n, err := file.Read(buffer)
	if err != nil {
		log.Fatalf("error reading from file: %v\n", err)
	}

	fmt.Printf("read %d bytes\n", n)
	for i, v := range buffer {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Printf("%s", string(buffer))
}
