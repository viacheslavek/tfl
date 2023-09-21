package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/parser"
)

func main() {

	example := "f(g(xxx), w(g(y), y)) -> g(f(x, y))\nh(g(x)) -> s(y)\ns(g(z), y) -> g(f(x, y))"

	expr := parser.InitExpression()

	if err := expr.ExtractPair(example); err != nil {
		fmt.Println("extract pair err", err)
	}

	if err := expr.ParseExpressionsToLinearRepresentation(); err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("_______________________________________________________")
	fmt.Println("_______________________________________________________")

	toBring := expr.EPs[0].Left
	bringExpr, err := expr.BringingLinearForm(toBring)
	fmt.Println("bring: ", bringExpr)
	if err != nil {
		fmt.Println("error bring:", err)
	}

	fmt.Println("END")

}
