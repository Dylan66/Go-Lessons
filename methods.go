package main

import "fmt"

type triangle struct {
	height int
	base int

}

func (tri triangle) area () float64 {
	return float64(tri.base*tri.height)*(0.5)

}

func main() {
	tri := triangle{height: 18, base: 20}
	fmt.Printf("The area of the triangle is %.2f", tri.area())
}
