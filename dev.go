package main

import "fmt"

func main() {

	messagesFromDoris :=[] string{
		"Get the Cake",
		"Help your Sister",
		"Cook the Dinner",
		"hustle hard",
			
	}
	numMessages :=float64(len(messagesFromDoris))
	costPerMessage :=0.02

	//this is the solution
	totalCost:= numMessages*costPerMessage
	fmt.Println(totalCost)
}