package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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








func main() {
	fmt.Println("Learning CRUD Operations in Go lang")
	// performGetRequest()
}
