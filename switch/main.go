package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	month := "January"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	case "July", "August", "September":
		fmt.Println("Summer")
	case "October", "November", "December":
		fmt.Println("Autumn")
	default:
		fmt.Println("Invalid month")
	}

	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 15:
		fmt.Println("Cold")
	case temperature >= 15 && temperature < 25:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}
