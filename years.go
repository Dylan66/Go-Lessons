//explicit returns, if statements and functions
package main

import "fmt"

func yearsToGoal(age int)(yearsDrink,yearsDrive,yearsRental int){
	yearsDrink=21-age
	if yearsDrink<0{
		yearsDrink=0

	}
	yearsDrive= 18-age
	if yearsDrive<0{
		yearsDrive=0
	}

	yearsRental=25-age
	if yearsRental<0 {
		yearsRental=0
	}
	return yearsDrink, yearsDrive, yearsRental

}
func testAge(age int) {
	fmt.Println("The age is:", age)
	yearsDrink, yearsDrive, yearsRental := yearsToGoal(age)

	fmt.Printf("You can drink in %d years, drive in %d years and rent in %d years/n," ,yearsDrink, yearsDrive, yearsRental)
}

	
func main () {
	testAge(12)

}