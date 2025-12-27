package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file:", err)
			return
		}
		defer file.Close()

		fmt.Println("File created successfully")

		content := "Hello, World!"
		bytes, error := io.WriteString(file, content)
		if error != nil {
			fmt.Println("Error while writing file", error)
			return
		}

		fmt.Println("Bytes:", bytes)
	*/

	file, error := os.Open("example.txt")
	if error != nil {
		fmt.Println("Error while opening file:", error)
		return
	}

	defer file.Close()

	/*
		// Create a buffer to read the file
		buffer := make([]byte, 1024)

		// Read the file contents into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file:", err)
				return
			}

			// Process the read content
			fmt.Println(string(buffer[:n]))
		}
	*/

	// Read the file using ioutil package
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading the file:", err)
		return
	}

	fmt.Println(string(content))

}
