//structs are key value pairs they are similar to dictionaries in python
//my first struct
/*
package main

import "fmt"

type carType struct{
    toyota string
    year int
    model string
    color string
    frontWheel wheel
    backWheel wheel

}
type wheel struct
{
    radius int
    material string
}
func checkCar(ccarType carType) bool {
    if ccarType.year==0 {
    return false
    }

    if ccarType.model==""{
        return false
    }
    if ccarType.frontWheel.radius==0 {
        return false
    }

    return true


}

//function to test the struct
func test( c carType){
    fmt.Println("The brand of the car is", c.toyota, "and its color is", c.color)
    fmt.Println("The model of the car is", c.model, "and its year is", c.year)
    fmt.Printf("The car's manufacture year is %d and its model is %s", c.year, c.model)
    fmt.Println("The front wheel radius is", c.frontWheel.radius, "and its material is", c.frontWheel.material)
}
func main(){
    car:= carType{"mark x", 0, "mark x evolution", "black", wheel{20, "rubber"}, wheel{20, "rubber"}}
   
   
    if checkCar(car) {
        fmt.Println("The car is valid")
        test(car)
    } else {
        fmt.Println("The car is not valid")
    }


}

*/
//embedded struct
package main

import "fmt"

type vehicle struct {
    model string
    year int
}

type lorry struct {
    vehicle
    weight int
    color string
}

func main() {
    bigCar := lorry{vehicle{"volvo", 2025}, 5000, "green"}
    fmt.Println("The model of the car is", bigCar.model)
    fmt.Println("The year of manufacture is", bigCar.year)
}
