package main

import "fmt"

func main() {

	// fmt.Println("==== Initially ====")
	// numbers := []int{1, 2, 3, 4, 5}
	// fmt.Println("Numbers are :", numbers)
	// fmt.Println("Length of Numbers is :", len(numbers))
	// fmt.Printf("Number has Data Type : %T\n", numbers)

	// // Appending values to slice
	// fmt.Println("==== After Appending ====")
	// numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	// fmt.Println("Numbers are :", numbers)
	// fmt.Println("Length of Numbers is :", len(numbers))
	// fmt.Printf("Number has Data Type : %T\n", numbers)

	// // Empty slice
	// fmt.Println("==== Empty Slice ====")
	// names := []string{}
	// fmt.Println("Names are :", names)
	// fmt.Println("Length of Names is :", len(names))
	// fmt.Printf("Names has Data Type : %T\n", names)



	// numbers := []int{1, 2, 3}
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))
	// fmt.Printf("Slice has Data Type : %T\n", numbers)



	numbers := make([]int, 3, 5)
	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	fmt.Printf("Slice has Data Type : %T\n", numbers)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	fmt.Printf("Slice has Data Type : %T\n", numbers)

	numbers = append(numbers, 6)
	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	fmt.Printf("Slice has Data Type : %T\n", numbers)

	

	// var stock = []string // Not allowed
	// stock = []string{} // Allowed
	stock := make([]string, 0) // Allowed
	fmt.Println("Slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))
	fmt.Printf("Slice has Data Type : %T\n", stock)
}
