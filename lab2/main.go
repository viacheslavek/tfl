package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/parser"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/reggen"
)

func main() {
	fmt.Println("start")

	regGenerator, _ := reggen.New(1, 5, 3, 15)

	regexes := regGenerator.Generate()

	fmt.Println("regex:", regexes[0])

	// regex := "((bc)*)*"
	regex := regexes[0]

	err := parser.ParseRegex(regex)
	if err != nil {
		fmt.Println("Проблема", err)
	}
}
