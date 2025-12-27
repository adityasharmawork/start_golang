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
	State string
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


	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Aditya",
		LastName: "Sharma",
		Age: 20,
	}

	employee1.Contact_Details.Email = "aditya@gmail.com"
	employee1.Contact_Details.Phone = "1234567890"

	employee1.Address_Details = Address{
		House: 123,
		Area: "Dhanbad",
		State: "Jharkhand",
	}

	fmt.Println("Employee 1:", employee1)

	// employee1 := Employee{
	// 	Person_Details: Person{
	// 		FirstName: "Aditya",
	// 		LastName: "Sharma",
	// 		Age: 20,
	// 	},
	// }
}
