package main

import "fmt"

func main() {
	fmt.Println("Learning Arrays in Go Lang")

	// var name[5]string

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Enter name for %d Person: ", i+1)
	// 	fmt.Scan(&name[i])
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("The name of", i+1, "Person is:", name[i])
	// }

	// fmt.Println("Enjoying learning Go Lang!!!")

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are:", numbers)

	fmt.Println("Length of numbers array is:", len(numbers))

	var price[5]int
	fmt.Println("Price is:", price)

	var names[5]string
	names[1] = "Aditya"
	fmt.Printf("Price is %q\n", names)
}
