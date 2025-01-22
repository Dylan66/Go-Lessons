//structs are key value pairs they are similar to dictionaries in python
//my first struct
package main

import "fmt"

type carType struct{
    toyota string
    year int
    model string
    color string

}
func test( c carType){
    fmt.Println("The brand of the car is", c.toyota, "and its color is", c.color)
    fmt.Println("The model of the car is", c.model, "and its year is", c.year)
    fmt.Printf("The car's manufacture year is %d and its model is %s", c.year, c.model)
}
func main(){
    test(carType{"mark x", 2019, "mark x evolution", "black"})


}