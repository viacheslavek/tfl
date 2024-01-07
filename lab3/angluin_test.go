package main

import (
	"fmt"
	"testing"

	"github.com/VyacheslavIsWorkingNow/tfl/lab3/tables"
)

// TODO: для теста нужно только сделать, чтобы оракул распознавал слова с b на позиции два с конца

func TestBinSecondFromEnd(t *testing.T) {
	testRun()
}

type handsOracleTest struct{}

func (ho handsOracleTest) BelongLanguage(word string) bool {
	var newAns string
	fmt.Println(word)
	_, _ = fmt.Scan(&newAns)
	return newAns == "yes"
}

func (ho handsOracleTest) GetAlphabet() []byte {
	return []byte{'a', 'b'}
}

func testRun() {
	ho := handsOracleTest{}
	a := tables.New(ho, 4)
	// оракул - вторая буква с конца - 'b'
	// первый этап - пока все пусто
	a.PrintTable()
	a.PrintExtendTable()
	fmt.Printf("first closed? '%s' -> yes?\n", a.Closed())
	fmt.Printf("second consistent? '%s' -> yes?\n", a.Consistent())

	// второй этап - из МАТа приходит 'ba'
	a.AddPrefix("ba")
	a.PrintTable()
	a.PrintExtendTable()

	fmt.Printf("second closed? '%s' -> yes?\n", a.Closed())
	fmt.Printf("second consistent? '%s' -> no?\n", a.Consistent())

	// третий этап - приходит суффикс 'а' из-за неконсистентности
	a.AddSuffix("a")
	a.PrintTable()
	a.PrintExtendTable()
	fmt.Printf("third closed? '%s' -> no?\n", a.Closed())
	fmt.Printf("second consistent? '%s' -> yes?\n", a.Consistent())

	// четвертый этап - из-за не закрытости приходит 'bb'
	a.AddPrefix("bb")
	a.PrintTable()
	a.PrintExtendTable()
	fmt.Printf("forth closed? '%s' -> yes?\n", a.Closed())
	fmt.Printf("second consistent? '%s' -> yes?\n", a.Consistent())
}
