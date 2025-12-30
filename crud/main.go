package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting response:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting response:", res.Status)
	}

	// fmt.Println("Response before IO Util Parsing:", res)

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response:", err)
	// 	return
	// }
	// fmt.Println("Data:", string(data))

	// Encoding to JSON (Marshalling)

	// var todo Todo

	// Decoding from JSON to Struct (Unmarshalling)

	// var todo Todo
	// err = json.Unmarshal(data, &todo)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Todo:", todo)

	// OR -

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Todo:", todo)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Aditya Sharma",
		Completed: true,
	}

	// Convert the Struct to JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while Marshalling:", err)
		return
	}

	jsonString := string(jsonData)

	fmt.Println("Data:", jsonString)

	// Convert JSON String to Reader

	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// Send POST Request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the Data:", err)
		return
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response:", string(data))

}

func performUpdateRequest() {	
	todo := Todo{
		UserID:    23789,
		Title:     "Aditya Sharma Go lang Hello World",
		Completed: true,
	}

	// Convert the Todo Struct to JSON Encoding (Marshal)
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while Marshalling Data:", err)
		return
	}


	// Create PUT Request
	// As there is no http.Put() method, we will use http.NewRequest()

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	const myURL = "https://jsonplaceholder.typicode.com/todos/1"


	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error while creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the PUT Request
	// For this, we will need to create a Client which will send the Request

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending the PUT Request from the Client:", err)
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the Data:", err)
		return
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response:", string(data))

}

func main() {
	fmt.Println("Learning CRUD Operations in Go lang")
	// performGetRequest()
	// performPostRequest()
	performUpdateRequest()
}
