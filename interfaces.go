package main

import "fmt"


type shapes interface {
	area() float64
	perimeter () float64
	color () string
}
 type rectangle struct {
	width float64
	length float64
 }

 func (rect rectangle) area() float64 {
	return rect.length*rect.width
 }
 func (rect rectangle) perimeter() float64 {
	return rect.length*2 + rect.width*2
 }
 func (rect rectangle) color() string {
	return "red"
 }

 type circle struct {
	// height represents the height of an object in float64 format.
	radius float64
 }
//this represents a circle
 func (cir circle) area() float64 {
	return 3.14*cir.radius*cir.radius
 }
func (cir circle) perimeter() float64 {
	return 3.14*cir.radius*2
}
func (cir circle) color() string {
	return "yellow"
}

func main () {
	rect := rectangle {width:10 ,length:50}
	fmt.Printf(" The area of the rectangle is %.2f and the perimeter is %.2f \n", rect.area(), rect.perimeter())
	fmt.Println("the color is", rect.color())

	cir := circle {radius : 21}
	fmt.Printf("the area of the circle is %.2f and perimeter is %.2f \n", cir.area(), cir.perimeter())
	fmt.Println("the color is", cir.color()) 


}