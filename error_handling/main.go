package main

import "fmt"

// func divide(a, b int) int {
// 	if(b == 0) {
// 		return 0
// 	}
// 	return a / b                              // 2
// }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil // 2.5
}

func main() {

	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	ans, _ := divide(10, 0)
	fmt.Println("Answer is ", ans)

}
