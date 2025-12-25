package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	fmt.Printf("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Printf("Hello, %s\n", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("Hello, Mr.", name)

}
