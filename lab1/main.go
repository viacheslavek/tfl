package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/parser"
)

func main() {

	example := "f(g(xxx), w(g(y), y)) -> g(f(x, y))\nh(g(x)) -> s(y)\ns(y) -> g(f(x, y))"

	expr := parser.InitExpression()

	if err := expr.ExtractPair(example); err != nil {
		fmt.Println("extract pair err", err)
		panic(err)
	}

	if err := expr.ParseExpressionsToLinearRepresentation(); err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	fmt.Println("_______________________________________________________")
	fmt.Println("_______________________________________________________")

	bringExpr, err := expr.BringingSuchForLinearForms()
	fmt.Println("bring: ", bringExpr)
	if err != nil {
		fmt.Println("error bring:", err)
	}

	fmt.Println("END")

}
