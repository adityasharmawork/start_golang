package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning Web Service!")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET Response:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println(res)
	fmt.Printf("The Type of Response: %T\n", res)

	// Read the response body

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error while reading response")
		return
	}

	fmt.Println("Response:", string(data))

}
