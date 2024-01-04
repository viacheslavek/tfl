package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"
)

func main() {
	fmt.Println("Hello")
	firstOracle := oracle.NewRegularOracle("aba*", []byte{'a', 'b'})

	fmt.Println(firstOracle.BelongLanguage("ab"))
	fmt.Println(firstOracle.BelongLanguage("aba"))
	fmt.Println(firstOracle.BelongLanguage("abaa"))
	fmt.Println(firstOracle.BelongLanguage("abaaaa"))
	fmt.Println(firstOracle.BelongLanguage("abaaaaa"))
	fmt.Println(firstOracle.BelongLanguage("abbb"))
	fmt.Println(firstOracle.BelongLanguage("abb"))

}

// TODO: в tables два файла для работы с разными таблицами
