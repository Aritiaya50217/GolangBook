package main

import "fmt"

type Engine struct{}

type Adder interface {
	Add(x, y float64) float64
}

func (e Engine) Add(x, y float64) float64 {
	return x + y
}

func NewEngine() *Engine {
	return &Engine{}
}

type Calculator struct {
	Adder Adder
}

func NewCalculator(a Adder) *Calculator {
	return &Calculator{Adder: a}
}

func InitCalc() *Calculator {
	adder := NewEngine()
	calculator := NewCalculator(adder)
	return calculator
}

func (c Calculator) PrintAdd(x, y float64) {
	fmt.Println("Result : ", c.Adder.Add(x, y))
}

func main() {
	calc := InitCalc()
	calc.PrintAdd(2.5, 3.5)
}
