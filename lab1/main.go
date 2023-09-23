package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/parser"
	"os"
	"os/exec"
)

// TODO: заменить паники на что-то более нежное

func main() {

	// example := "f(g(x), w(g(y), z)) -> g(f(x, y))\nh(g(x)) -> s(y)\ns(y) -> g(f(x, y))"

	// На вход подать неравенства через стрелку и знак \n

	example := "h(x, y) -> s(x, y)"

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

	cmd := exec.Command("z3", "lab1/solver.smt2")

	output, eErr := cmd.CombinedOutput()
	if wErr != nil {
		panic(eErr)
	}

	fmt.Println("Результат выполнения команды:")
	fmt.Println(string(output))

}
