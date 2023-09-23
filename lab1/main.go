package main

import (
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/parser"
	"os"
)

func main() {

	// example := "f(g(x), w(g(y), z)) -> g(f(x, y))\nh(g(x)) -> s(y)\ns(y) -> g(f(x, y))"

	example := "h(x) -> s(x)"

	report, err := parser.Parse(example)
	if err != nil {
		panic(err)
	}

	file, oErr := os.Create("lab1/solver.smt2")
	if oErr != nil {
		panic(oErr)
	}
	defer func() {
		_ = file.Close()
	}()

	_, wErr := file.WriteString(report)
	if wErr != nil {
		panic(wErr)
	}

}
