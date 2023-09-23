package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/parser"
)

func main() {

	// example := "f(g(x), w(g(y), z)) -> g(f(x, y))\nh(g(x)) -> s(y)\ns(y) -> g(f(x, y))"

	example := "h(g(x)) -> s(y)\ns(y) -> g(f(x, y))"

	expr := parser.InitExpression()

	if err := expr.ExtractPair(example); err != nil {
		fmt.Println("extract pair err", err)
		panic(err)
	}

	if err := expr.ParseExpressionsToLinearRepresentation(); err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	err := expr.BringingSuchForLinearForms()
	if err != nil {
		fmt.Println("error bring:", err)
		panic(err)
	}

	fmt.Println(expr.ToStringConstructions())

	fmt.Println("END")

	fmt.Println(expr.MakeFirstInequality())

	fmt.Println(expr.MakeSecondInequality())

}
