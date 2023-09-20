package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/parser"
)

// TODO: добавить логи, чтобы лучше понимать, что к чему и где
// TODO: ошибки в err_name

func main() {

	example := "f(g(xxx), ttt) -> g(f(x, y))\nh(g(x)) -> s(y)\ns(g(z), y) -> g(f(x, y))"

	expr := parser.InitExpression()

	if err := expr.ExtractPair(example); err != nil {
		fmt.Println("extract pair err", err)
	}

	myNewPair := expr.EPs[0]

	fmt.Println("left:", myNewPair.Left)

	temp, err := expr.ParseExpressionsToLinearRepresentation()
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("repr:", temp)

	fmt.Println(expr.Variables)

	fmt.Println("END")

}
