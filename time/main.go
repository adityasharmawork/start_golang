package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("Type of currentTime is %T\n", currentTime)

	formatted := currentTime.Format("2006/01/02, 03:04:05 PM")
	fmt.Println(formatted)

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println(formatted_time)

	// add one more day in current time
	new_time := currentTime.Add(24 * time.Hour)
	fmt.Println("New time is:", new_time)
	formatted_new_date := new_time.Format("2006/01/02 Monday")
	fmt.Println(formatted_new_date)
}
