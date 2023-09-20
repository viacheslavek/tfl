package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/parser"
)

func main() {

	example := "f(g(xxx), ttt) -> g(f(x, y))\nh(g(x)) -> s(y)\ns(g(z), y) -> g(f(x, y))"

	expr := parser.InitExpression()

	if err := expr.ExtractPair(example); err != nil {
		fmt.Println("extract pair err", err)
	}

	if err := expr.ParseExpressionsToLinearRepresentation(); err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("END")

}
