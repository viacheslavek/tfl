package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/wordgen"

	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/gluskov"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/loop"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/parser"
)

func main() {
	fmt.Println("start")

	// TODO: сканирую параметры для генерации регулярок
	// переношу в reggen.New()
	//regGenerator, _ := reggen.New(3, 5, 3, 15)
	//regexes := regGenerator.Generate()
	//

	regex := "(a(abc)*)b*"

	// TODO: можно сканировать параметры для генерации слов в регулярках (максимальное число накачки и число слов)

	_ = parser.ParseRegexInDot(regex)

	tree, err := parser.ParseRegex(regex)

	if err != nil {
		fmt.Println("беда в парсере", err)
	}

	automaton := gluskov.BuildMachine(tree)

	err = automaton.GetDotMachine()
	if err != nil {
		fmt.Println("беда с визуализацией автомата", err)
	}

	loops := loop.FindCycles(automaton)
	letterLoop := loop.TranslateLoops(loops, automaton)
	fmt.Printf("letter loop: %+v\n", letterLoop)

	org := wordgen.New(regex, 10, 5)

	org.DfsBuildWord(automaton, letterLoop)

}
