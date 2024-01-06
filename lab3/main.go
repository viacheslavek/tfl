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
	fmt.Println("Hello")
	alphabet := []byte{'a', 'b'}
	firstOracle := oracle.NewRegularOracle("aba*", alphabet)

	fmt.Println(firstOracle.BelongLanguage("ab"))
	fmt.Println(firstOracle.BelongLanguage("abb"))

	fmt.Println(firstOracle.BelongLanguage("abaaaa"))

	ho := handsOracle{}

	angluin := tables.New(ho)

	fmt.Printf("simple angluin %+v\n", angluin)

	angluin.Run()

	fmt.Println("SUCCESS")

}

// TODO: в tables два файла для работы с разными таблицами
