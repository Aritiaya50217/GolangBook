package main

import "fmt"

// IDrawShape interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

// DrawShape struct
type DrawShape struct{}

// DrawShape struct has method draw Shape with float x and y coordinates
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Printf("Drawing Shape")
}

// IContour interface
type IContour interface {
	drawCoutour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

// DrawCoutour struct
type DrawCoutour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

// DrawContour method drawContour given the coordinates
func (contour DrawCoutour) drawCoutour(x [5]float32, y [5]float32) {
	fmt.Println("Draw Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

// DrawContour method resizeByFactor given factor
func (contour DrawCoutour) resizeByFactor(factor int) {
	contour.factor = factor
}
func main() {
	defaults := [5]float32{1, 2, 3, 4, 5}
	x, y := defaults, defaults
	var contour IContour = DrawCoutour{x, y, DrawShape{}, 2}
	fmt.Println("countour : ", contour)
	contour.drawCoutour(x, y)
	contour.resizeByFactor(2)

}
