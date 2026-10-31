package main

import (
	"flag"
	"log"
	"test-driven-development/part01the_big_picture/chapter03mocking_and_assertion_frameworks/02exploring_mocks/calculator"
	"test-driven-development/part01the_big_picture/chapter03mocking_and_assertion_frameworks/02exploring_mocks/input"
)

func main() {
	expr := flag.String("expression", "", "mathematical expression to parse")
	flag.Parse()

	engine := calculator.NewEngine()
	validator := input.NewValidator(engine.GetNumOperands(), engine.GetValidOperators())
	parser := input.NewParser(engine, validator)
	result, err := parser.ProcessExpression(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*result)
}
