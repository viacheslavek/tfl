package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"
	"github.com/VyacheslavIsWorkingNow/tfl/lab3/tables"
)

func main() {
	fmt.Println("Hello")
	alphabet := []byte{'a', 'b'}
	firstOracle := oracle.NewRegularOracle("aba*", alphabet)

	fmt.Println(firstOracle.BelongLanguage("ab"))
	fmt.Println(firstOracle.BelongLanguage("abb"))

	angluin := tables.New(firstOracle)

	fmt.Printf("simple angluin %+v\n", angluin)

	angluin.Run()

	fmt.Println("SUCCESS")

}

// TODO: в tables два файла для работы с разными таблицами
