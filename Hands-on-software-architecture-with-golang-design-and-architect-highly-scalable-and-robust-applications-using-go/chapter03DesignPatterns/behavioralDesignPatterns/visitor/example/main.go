package main

import "fmt"

// Element
type Shape interface {
	Accept(v Visitor)
}

// concrete elements
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitorCircle(c)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitorRectangle(r)
}

// Visitor
type Visitor interface {
	VisitorCircle(c *Circle)
	VisitorRectangle(r *Rectangle)
}

// concrete visitor
type AreaCalculator struct{}

func (a *AreaCalculator) VisitorCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	fmt.Printf("Circle Area : %.2f\n", area)
}

func (a *AreaCalculator) VisitorRectangle(r *Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("Rectangle Area: %.2f\n", area)
}

type SummaryPrinter struct{}

func (s *SummaryPrinter) VisitorCircle(c *Circle) {
	fmt.Printf("Circle: radius = %.2f\n", c.Radius)
}

func (s *SummaryPrinter) VisitorRectangle(r *Rectangle) {
	fmt.Printf("Rectangle: width = %.2f, height = %.2f\n", r.Width, r.Height)
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Rectangle{Width: 3, Height: 4},
	}

	fmt.Println("------ Area Calculation ------")
	areaVisitor := &AreaCalculator{}
	for _, shape := range shapes {
		shape.Accept(areaVisitor)
	}

	fmt.Println("\n------ Summary Printing ------")
	printVisitor := &SummaryPrinter{}
	for _, shape := range shapes {
		shape.Accept(printVisitor)
	}
}
