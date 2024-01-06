package main

import (
	"fmt"
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
	//alphabet := []byte{'a', 'b'}
	//firstOracle := oracle.NewRegularOracle("aba*", alphabet)
	//
	//fmt.Println(firstOracle.BelongLanguage("ab"))
	//fmt.Println(firstOracle.BelongLanguage("abb"))
	//
	//fmt.Println(firstOracle.BelongLanguage("abaaaa"))
	//
	//ho := handsOracle{}
	//
	//angluin := tables.New(ho)
	//
	//fmt.Printf("simple angluin %+v\n", angluin)
	// angluin.Run()

	fmt.Println("SUCCESS")

}
