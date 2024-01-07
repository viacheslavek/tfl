package main

import (
	"testing"

	"github.com/VyacheslavIsWorkingNow/tfl/lab3/tables"
)

type BinSecondFromEnd struct{}

func (b BinSecondFromEnd) BelongLanguage(word string) bool {
	if len(word) < 2 {
		return false
	}
	return word[len(word)-2] == 'b'
}

func (b BinSecondFromEnd) GetAlphabet() []byte {
	return []byte{'a', 'b'}
}

func TestBinSecondFromEnd(t *testing.T) {
	bsfe := BinSecondFromEnd{}

	angluin := tables.New(bsfe, 4)
	auto := angluin.Run()
	err := auto.GetDotMachine("testOracleB")
	if err != nil {
		t.Errorf("Ошибка визуализации в тесте: %e", err)
	}
}
