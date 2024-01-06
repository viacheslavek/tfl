package tables

import (
	"fmt"
	"log"

	"github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"
)

type Angluin struct {
	suffix       map[string]struct{}
	prefix       map[string]struct{}
	extendPrefix map[string]struct{}
	table        map[string]bool
	extendTable  map[string]bool
	oracle       oracle.Oracle
}

func New(o oracle.Oracle) *Angluin {
	a := Angluin{
		suffix:       make(map[string]struct{}),
		prefix:       make(map[string]struct{}),
		extendPrefix: make(map[string]struct{}),
		table:        make(map[string]bool),
		extendTable:  make(map[string]bool),
		oracle:       o,
	}

	a.table["_"] = a.oracle.BelongLanguage("")

	a.prefix[""] = struct{}{}
	a.suffix[""] = struct{}{}

	a.addExtendPrefix("")

	return &a
}

// Run TODO: Это делаю уже в последнюю очередь:
// Запускаю прогон, если таблица консистентна и полна, то кидаю в учителя
// если все норм, то отдаю автомат, если нет, то c новой строкой повторяю итерацию
func (a *Angluin) Run() {
	fmt.Println("in RUN")

	a.testRun()

	fmt.Println("end in run")
}

// Closed INFO: Closed: An observation table is called closed if for all t in S.A there exist an s` in S
// such that row(s’)=row(t).This states that every row(s.a) must be present in row(s).
// Если полна, то вернется пустая строка, если нет, то префикс
func (a *Angluin) Closed() string {

	suffixList := sortSet(a.suffix)
	tempTableRowMap := a.createTempTableRowMap(suffixList)
	extendPrefixList := sortSet(a.extendPrefix)
	for _, ep := range extendPrefixList {
		if _, ok := tempTableRowMap[a.getExtendTableRow(ep, suffixList)]; !ok {
			return ep
		}
	}

	return ""
}

func (a *Angluin) createTempTableRowMap(suffixList []string) map[string]struct{} {
	tempTableRowMap := make(map[string]struct{})

	for p := range a.prefix {
		tempTableRowMap[a.getTableRow(p, suffixList)] = struct{}{}
	}

	return tempTableRowMap
}

func (a *Angluin) getTableRow(prefix string, suffixList []string) string {
	row := ""
	for _, s := range suffixList {
		val := a.table[createTableKey(prefix, s)]
		if val {
			row += "1"
		} else {
			row += "0"
		}
	}
	return row
}

func (a *Angluin) getExtendTableRow(prefix string, suffixList []string) string {
	row := ""
	for _, s := range suffixList {
		val := a.extendTable[createTableKey(prefix, s)]
		if val {
			row += "1"
		} else {
			row += "0"
		}
	}
	return row
}

// Consistent INFO: Consistent: An observation table is said to be consistent if, whenever s1,s2 in S satisfy row(s1)=row(s2)
// then for every an in A must satisfy row(s1.a)=row(s2.a).
// Если консистентно, то вернется пустая строка, иначе - буква + суффикс
func (a *Angluin) Consistent() string {

	suffixList := sortSet(a.suffix)

	tablePrefixToRow := a.getDsForTablePrefixToRow(suffixList)
	extendTablePrefixToRow := a.getDsForExtendTablePrefixToRow(suffixList)

	equalTableRowToPrefix := getEqualRowForPrefix(tablePrefixToRow)

	return a.findConsistentForRowInTables(equalTableRowToPrefix, tablePrefixToRow, extendTablePrefixToRow, suffixList)
}

func (a *Angluin) getDsForTablePrefixToRow(suffixList []string) map[string]string {
	tableRowToPrefix := make(map[string]string)

	for p := range a.prefix {
		tableRowToPrefix[p] = a.getTableRow(p, suffixList)
	}

	return tableRowToPrefix
}

func (a *Angluin) getDsForExtendTablePrefixToRow(suffixList []string) map[string]string {
	extendTableRowToPrefix := make(map[string]string)

	for ep := range a.extendPrefix {
		extendTableRowToPrefix[ep] = a.getExtendTableRow(ep, suffixList)
	}

	return extendTableRowToPrefix
}

func getEqualRowForPrefix(prefixToRow map[string]string) map[string][]string {
	result := make(map[string][]string)
	for prefix, row := range prefixToRow {
		result[row] = append(result[row], prefix)
	}

	return result
}

func (a *Angluin) findConsistentForRowInTables(
	equalTableRowToPrefix map[string][]string,
	tablePrefixToRow, extendTablePrefixToRow map[string]string,
	suffixList []string) string {

	for _, prefixes := range equalTableRowToPrefix {
		for i, prefix1 := range prefixes {
			for _, prefix2 := range prefixes[i+1:] {
				for _, letter := range a.oracle.GetAlphabet() {
					combined1 := getWordWithLetter(prefix1, letter)
					combined2 := getWordWithLetter(prefix2, letter)
					if ok, row1, row2 := a.consistentForPair(
						combined1, combined2, tablePrefixToRow, extendTablePrefixToRow,
					); !ok {
						return findSuffixAndLetterInRow(row1, row2, letter, suffixList)
					}
				}
			}
		}
	}

	return ""
}

func getWordWithLetter(p string, l byte) string {
	return p + string(l)
}

func (a *Angluin) consistentForPair(
	prefixLetter1, prefixLetter2 string, tablePrefixToRow, extendTablePrefixToRow map[string]string,
) (bool, string, string) {

	var row1, row2 string
	var ok1, ok2 bool

	row1, ok1 = tablePrefixToRow[prefixLetter1]
	if !ok1 {
		row1, ok1 = extendTablePrefixToRow[prefixLetter1]
		if !ok1 {
			log.Println("row1", prefixLetter1, "not find")
		}
	}

	row2, ok2 = tablePrefixToRow[prefixLetter2]
	if !ok2 {
		row2, ok2 = extendTablePrefixToRow[prefixLetter2]
		if !ok2 {
			log.Println("row2", prefixLetter2, "not find")
		}
	}

	return row1 == row2, row1, row2
}

func findSuffixAndLetterInRow(row1, row2 string, letter byte, suffixList []string) string {

	for i := 0; i < len(row1); i++ {
		if row1[i] != row2[i] {
			return string(letter) + suffixList[i]
		}
	}

	log.Println("This is can not happened. row1 == row2 but row1 != row2")

	return ""
}

func (a *Angluin) testRun() {
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
	fmt.Println("ex:", a.extendTable)

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
