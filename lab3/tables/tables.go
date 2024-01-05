package tables

import (
	"fmt"
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
	return &a
}

// Run TODO: Это делаю уже в последнюю очередь:
// Запускаю прогон, если таблица констистента и полна, то кидаю в учителя
// если все норм, то отдаю автомат, если нет, то c новой строкой повторяю итерацию
func (a *Angluin) Run() {
	fmt.Println("in RUN")

	a.AddPrefix("a")
	a.AddPrefix("aa")
	a.PrintTable()

	a.AddSuffix("b")
	a.PrintTable()

	a.AddSuffix("ba")
	a.PrintTable()

	a.AddExtendPrefix("a")
	a.AddSuffix("aaa")

	a.PrintPrefix()
	a.PrintSuffix()
	a.PrintExtendPrefix()
	a.PrintTable()
	a.PrintExtendTable()

	fmt.Println("end in run")
}

// INFO:Closed: An observation table is called closed if for all t in S.A there exist an s’ in S
// such that row(s’)=row(t).This states that every row(s.a) must be present in row(s).
// Если полна, то вернется пустая строка, если нет, то префикс
func (a *Angluin) closed() string {

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

// INFO: Consistent: An observation table is said to be consistent if, whenever s1,s2 in S satisfy row(s1)=row(s2)
// then for every a in A it must satisfy row(s1.a)=row(s2.a).
// TODO: проверяю констистентость как в презе
func (a *Angluin) consistent() {

}
