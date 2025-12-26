package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area string
	state string
}

type Employee struct {
	Person_Details Person
	Contact_Details Contact
	Address_Details Address
}

func main() {
	// var aditya Person
	// aditya.FirstName = "Aditya"
	// aditya.LastName = "Sharma"
	// aditya.Age = 20
	// fmt.Println(aditya)

	person1 := Person{
		FirstName: "Harkirat",
		LastName: "Singh",
		Age: 20,
	}

	fmt.Println(person1)


	// new keyword

	// var person2 = new(Person)
	// person2.FirstName = "Harkirat"
	// person2.LastName = "Singh"
	// person2.Age = 20
	// fmt.Println(person2)


	employee1 := Employee{
		Person_Details: Person{
			FirstName: "Aditya",
			LastName: "Sharma",
			Age: 20,
		},
	}
}
