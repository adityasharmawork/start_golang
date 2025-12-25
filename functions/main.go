package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

func add(a, b int) (int) {
	return a + b
}

func main() {
	fmt.Println("Hey there!")
	simpleFunction()

	ans := add(2, 3)
	fmt.Println("Sum of two numbers is :", ans)
}
