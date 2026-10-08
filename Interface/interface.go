package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Modifiers interface {
	Scale(float64)
}

type Geometry interface {
	Shape
	Modifiers
}

/**
 * Rectangle represents a rectangle with a width and height.
 * It will now be part of the Geometry interface since it 
 * implements both Shape and Modifiers.
 */

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

/**
 * Circle represents a circle with a radius.
 * It will now be part of the Geometry interface since it 
 * implements both Shape and Modifiers.
 */

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Scale(factor float64) {
	c.radius *= factor
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func ScaleShape(s Modifiers, factor float64) {
	s.Scale(factor)
}

func describeShape(g Geometry){
	fmt.Printf("Area: %.2f\n", g.Area())
	fmt.Printf("Perimeter: %.2f\n", g.Perimeter())
}

func main() {
	fmt.Println("Interface Example")
	rect := Rectangle{width: 3, height: 4}
	circle := Circle{radius: 5}

	describeShape(&rect)
	describeShape(&circle)

	fmt.Println("Scaling shapes...")
	ScaleShape(&rect, 2)
	ScaleShape(&circle, 0.5)

	describeShape(&rect)
	describeShape(&circle)


	/* 
		We can also use the Geometry interface to work with any shape 
		that implements both Shape and Modifiers.
		Se we can assing to a Geometry Interface any 
		shape that implements both Shape and Modifiers.(that in this case is a
		struct)
	*/
	var test Geometry = &Rectangle{width: 1, height: 2}
	fmt.Printf("Test Area: %.2f\n", test.Area())
	fmt.Printf("Test Perimeter: %.2f\n", test.Perimeter())
}