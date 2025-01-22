/*
package main

import "fmt"
// functions
func add(peter, lois int) int{
	return peter - lois

}
 func main () {
	averageOpenRate, displayMessage := .23, " is the average messagerate"
	fmt.Println("The credentials are:", averageOpenRate,displayMessage)
	averageAge:=2.5
	ageInt:=int(averageAge)
	fmt.Println("The age is,", ageInt)

	//constant  variables
	const minutesHour=60
	const secondsMinutes=60
	const totalTime=minutesHour*secondsMinutes
	fmt.Println("The total time is:", totalTime)
	
	//String formatting
	msg := fmt.Sprintf("the %v is the number of messages", 45)
	message :=fmt.Sprintf("The message is %s and the code is %d let the boys %v", "good",54, "cook")
	numberText:=fmt.Sprintf("The average number of messages is %.2f which is high", 84.304)
	fmt.Println(msg, message)
	fmt.Println(numberText)

	fmt.Println(concat("The name is,", "Jane"))
	fmt.Println(add(5,3))

}
func concat(piece string, piece2 string) string{
	return piece + piece2
}

*/
//
/*
package main

import "fmt"

func main () {
	messages := 60
	const messagesSent=40
	messages= incrementSends(messages, messagesSent)
	fmt.Println("You have sent", messages)
}
func incrementSends(messages, messagesSent int) int {
	return messages + messagesSent
}
*/


//returning multiple return values
package main 

import "fmt"

func main () {
	getNames():=firstName, _
	fmt.Println("The namee is", firstName)

}

func getNames() (string, string){
	return "Mary","Mangoes"
}