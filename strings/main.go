package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")

	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str = "    Hello, Go!    "
	fmt.Println(str)
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "Aditya"
	str2 := "Sharma"

	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result)

}
