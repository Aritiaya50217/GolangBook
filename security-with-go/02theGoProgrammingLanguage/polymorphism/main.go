package main

import "fmt"

type Number interface {
	int | float64
}

func add[T Number](a, b T) T {
	return a + b
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Redius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Redius * c.Redius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// ฟังก์ชันที่ใช้ polymorphism
func printArea(s Shape) {
	fmt.Println("Area : ", s.Area())
}

func main() {
	// compile-time
	fmt.Println(add(5, 10))
	fmt.Println(add(2.4, 3.6))

	// Run-time Polymorphism
	c := Circle{Redius: 5}
	r := Rectangle{Width: 4, Height: 6}
	// polymorphism: ใช้ function เดียวกัน แต่ทำงานต่างกัน
	printArea(c)
	printArea(r)
}
