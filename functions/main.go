package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

// func add(a, b int) int {
// 	return a + b
// }

func add(a, b int) (result int) {
	result = a + b
	return
}

func multiply(a, b int) (result int)  {
	result = a * b
	return
}

func main() {
	fmt.Println("Hey there!")
	simpleFunction()

	ans := add(2, 3)
	fmt.Println("Sum of two numbers is :", ans)

	data := multiply(2, 3)
	fmt.Println("Multiplication of two numbers is :", data)
}
