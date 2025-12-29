package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("Learning JSON")
	person := Person{
		Name:    "Aditya",
		Age:     20,
		IsAdult: true,
	}

	fmt.Println("Person details:", person)



	// Converting person to JSON Encoding (Marshalling)
	jsonPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error while marshalling:", err)
		return
	}
	fmt.Println("JSON Person:", string(jsonPerson))
	


	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonPerson, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	fmt.Println("Person Data:", personData)


}
