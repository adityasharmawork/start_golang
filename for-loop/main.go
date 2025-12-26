package main

import "fmt"

func main() {
	// for i := 0; i < 4; i++ {
	// 	fmt.Println("The number is:", i)
	// }

	// counter := 0

	// for {
	// 	fmt.Println("Infinite loop.")
	// 	counter++
	// 	if counter == 3 {
	// 		break
	// 	}
	// }

	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf("The value of the number at index %d is: %d\n", index, value)
	// }

	data := "Hello, World!"
	for index, char := range data {
		fmt.Printf("The character at index %d is: %c\n", index, char)
	}
}
