package main

import (
	"fmt"

	"github.com/VyacheslavIsWorkingNow/tfl/lab3/mat"
	"github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"
)

type handsOracle struct{}

func (ho handsOracle) BelongLanguage(word string) bool {
	var newAns string
	fmt.Println(word)
	fmt.Scan(&newAns)
	return newAns == "yes"
}

func (ho handsOracle) GetAlphabet() []byte {
	return []byte{'a', 'b'}
}

func main() {
	alphabet := []byte{'a', 'b'}
	firstOracle := oracle.NewRegularOracle("aba*", alphabet)

	//ho := handsOracle{}
	//
	//angluin := tables.New(ho, 4)
	//
	//fmt.Printf("simple angluin %+v\n", angluin)
	//angluin.Run()

	m := mat.New(firstOracle, 4)

	words := m.GenerateWords()

	fmt.Println("words:", words)

	fmt.Println("SUCCESS")

}
