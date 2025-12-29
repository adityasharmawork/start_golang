package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	fmt.Printf("The type of URL is: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("Error while parsing URL:", err)
		return
	}
	

	fmt.Printf("The type of Parsed URL is: %T\n", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)

	// Modifying URL Components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=aditya"

	// Constructing a URL string from a URL object
	newURL := parsedURL.String()
	fmt.Println("New URL:", newURL)

}
