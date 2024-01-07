package main

import (
	"fmt"

	"github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"
	"github.com/VyacheslavIsWorkingNow/tfl/lab3/tables"
)

func main() {

	fmt.Println("START")

	// TODO: добавить бинарный оракул и именно с ним будет запускаться main

	alphabet := []byte{'a', 'b'}
	regOracle := oracle.NewRegularOracle("aba*", alphabet)

	angluin := tables.New(regOracle, 4)
	angluin.Run()

	fmt.Println("SUCCESS")

}
