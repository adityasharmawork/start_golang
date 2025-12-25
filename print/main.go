package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.8234567

	fmt.Println("Name:", name, "Age:", age, "Height:", height) // It adds speces in between of different parts of the statement
	fmt.Println("Hello World")

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Height is %.2f\n", height)
	fmt.Printf("Type of name is %T\n", name)
}
