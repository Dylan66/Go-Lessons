package main

import "fmt"

func getDay(dayToday int) string {
	if dayToday ==1{
		return "mondays"
	}
	if dayToday==2{
		return "tuesdays"
	}
	if dayToday==3 {
		return "wednesdays"
	}
	return "unknown"
}

func main() {
	dayToday :=6
	day := getDay(dayToday)
	fmt.Println("Today is:", day)
}
