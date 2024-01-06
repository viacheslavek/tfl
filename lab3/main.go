package main

import (
	"fmt"

	"github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"
	"github.com/VyacheslavIsWorkingNow/tfl/lab3/tables"
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

	fmt.Println("START")

	alphabet := []byte{'a', 'b'}
	firstOracle := oracle.NewRegularOracle("aba*", alphabet)
	firstOracle.BelongLanguage("ab")

	ho := handsOracle{}
	angluin := tables.New(ho, 4)
	fmt.Printf("simple angluin %+v\n", angluin)
	angluin.Run()

	fmt.Println("SUCCESS")

}
