package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	// time.Sleep(2000 * time.Millisecond) // Simulating some work
	fmt.Println("Hello function completed successfully")
}

func sayHi() {
	fmt.Println("Hi Aditya :)")
	// time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Learning Goroutines")
	go sayHello()
	go sayHi()

	time.Sleep(1000 * time.Millisecond)
}
