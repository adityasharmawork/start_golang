package main

import "fmt"

func main() {

	fmt.Println("==== Initially ====")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are :", numbers)
	fmt.Println("Length of Numbers is :", len(numbers))
	fmt.Printf("Number has Data Type : %T\n", numbers)

	// Appending values to slice
	fmt.Println("==== After Appending ====")
	numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println("Numbers are :", numbers)
	fmt.Println("Length of Numbers is :", len(numbers))
	fmt.Printf("Number has Data Type : %T\n", numbers)

	// Empty slice
	fmt.Println("==== Empty Slice ====")
	names := []string{}
	fmt.Println("Names are :", names)
	fmt.Println("Length of Names is :", len(names))
	fmt.Printf("Names has Data Type : %T\n", names)
}
