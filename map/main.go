package main

import "fmt"

func main() {
	studentGrades := make(map[string]int) // name -> grade

	studentGrades["Aditya"] = 95
	studentGrades["Rahul"] = 85
	studentGrades["Ravi"] = 90

	fmt.Println(studentGrades)

	fmt.Println("Grade of Aditya is:", studentGrades["Aditya"])
	fmt.Println("Grade of Rahul is:", studentGrades["Rahul"])
	fmt.Println("Grade of Ravi is:", studentGrades["Ravi"])

	studentGrades["Ravi"] = 92
	fmt.Println("Grade of Aditya is:", studentGrades["Aditya"])
	fmt.Println("Grade of Rahul is:", studentGrades["Rahul"])
	fmt.Println("Grade of Ravi is:", studentGrades["Ravi"])

	delete(studentGrades, "Rahul")
	fmt.Println(studentGrades)
	fmt.Println("Grade of Aditya is:", studentGrades["Aditya"])
	fmt.Println("Grade of Rahul is:", studentGrades["Rahul"]) // 0
	fmt.Println("Grade of Ravi is:", studentGrades["Ravi"])


	grades, exists := studentGrades["David"]
	fmt.Println("Grades of David:", grades)
	fmt.Println("Exists:", exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is: %d\n", index, value)
	}
}
