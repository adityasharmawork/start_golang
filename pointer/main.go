package main

import "fmt"

func modifyValueByReference(val *int) {
	*val = 20
}

func main() {

	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num

	fmt.Println("Num has value:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Pointer has value:", ptr)
	fmt.Println("Address of pointer:", &ptr)
	fmt.Println("Value at pointer:", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is nil")
	}

	value := 10
	fmt.Println("Value before modification:", value)
	modifyValueByReference(&value)
	fmt.Println("Value after modification:", value)
}
