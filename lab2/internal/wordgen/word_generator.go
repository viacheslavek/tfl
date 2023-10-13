package wordgen

import (
	"fmt"

	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/gluskov"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/loop"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/parser"
)

type ManyRegexpWordGenerator struct {
	countRegex int
	rg         []OneRegexpGenerator
}

type OneRegexpGenerator struct {
	Regex       string
	CountWord   int
	MaxDumpSize int
	Words       []string
}

func New(regex string, countWord, maxDumpSize int) *OneRegexpGenerator {
	return &OneRegexpGenerator{
		Words:       make([]string, countWord),
		CountWord:   countWord,
		MaxDumpSize: maxDumpSize,
		Regex:       regex,
	}
}

func GenerateWordsForRegex(regex string, countWords, maxDumpSize int) error {
	org := New(regex, countWords, maxDumpSize)

	tree, pErr := parser.ParseRegex(regex)
	if pErr != nil {
		return fmt.Errorf("can't parse regex %w", pErr)
	}

	automaton := gluskov.BuildMachine(tree)

	loops := loop.FindCycles(automaton)
	letterLoop := loop.TranslateLoops(loops, automaton)

	for i := 0; i < countWords; i++ {
		org.Words[i] = org.DfsBuildWord(automaton, letterLoop)
	}

	return nil
}

// DfsBuildWord TODO: сделать dfs
func (org *OneRegexpGenerator) DfsBuildWord(
	machine *gluskov.Machine,
	loops loop.StateLoopToString,
) string {

	fmt.Println("start build word")

	fmt.Println(machine)
	fmt.Println(loops)

	return ""
}
