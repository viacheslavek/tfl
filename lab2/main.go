package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/parser"
)

func main() {
	fmt.Println("start")

	regex := "a*b*c|a"

	parser.ParseRegex(regex)
}
