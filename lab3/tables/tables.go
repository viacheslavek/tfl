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

	a.addExtendPrefix("")

	return &a
}

// Run TODO: Это делаю уже в последнюю очередь:
// Запускаю прогон, если таблица констистента и полна, то кидаю в учителя
// если все норм, то отдаю автомат, если нет, то c новой строкой повторяю итерацию
// TODO: перенести этот RUN в тест работы таблицы - оракул - вторая буква с конца - 'b'
func (a *Angluin) Run() {
	fmt.Println("in RUN")

	// первый этап - пока все пусто
	a.PrintTable()
	a.PrintExtendTable()
	fmt.Printf("first closed? '%s' -> yes?\n", a.Closed())

	// второй этап - из МАТа приходит 'ba'
	a.AddPrefix("ba")
	a.PrintTable()
	a.PrintExtendTable()
	fmt.Println("ex:", a.extendTable)

	fmt.Printf("second closed? '%s' -> yes?\n", a.Closed())

	// третий этап - приходит суффикс 'а' из-за неконсистентности
	a.AddSuffix("a")
	a.PrintTable()
	a.PrintExtendTable()
	fmt.Printf("third closed? '%s' -> no?\n", a.Closed())

	// четвертый этап - из-за не закрытости приходит 'bb'
	a.AddPrefix("bb")
	a.PrintTable()
	a.PrintExtendTable()
	fmt.Printf("forth closed? '%s' -> yes?\n", a.Closed())

	// после этого можно строить автомат

	fmt.Println("end in run")
}

// Closed INFO:Closed: An observation table is called closed if for all t in S.A there exist an s’ in S
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

	tableRowToPrefix := getDsForTableRowToPrefix()
	extendTableRowToPrefix := getDsForTableRowToPrefix()

	equalTableRowToPrefix := getEqualRowForPrefix(tableRowToPrefix)

	return a.findConsistentForRowInTables(equalTableRowToPrefix, tableRowToPrefix, extendTableRowToPrefix)
}

// TODO: считаю row для table вида prefix - row
func getDsForTableRowToPrefix() map[string]string {
	panic("Implement me")
}

// TODO: считаю row для extendTable вида prefix - row
func getDsForExtendTableRowToPrefix() map[string]string {
	panic("Implement me")
}

// TODO: в prefix - row нахожу такие пары, что row(prefix1) = row(prefix2)
// и создаю мапу вида row -> []prefix
func getEqualRowForPrefix(rowToPrefix map[string]string) map[string][]string {
	panic("Implement me")
}

// TODO: для всех row -> []prefix в []prefix попарно сопоставляю с алфавитом
// и ищу этот newPrefix в table или extendTable
func (a *Angluin) findConsistentForRowInTables(
	equalTableRowToPrefix map[string][]string, tableRowToPrefix, extendTableRowToPrefix map[string]string) string {

	// TODO: Если равны, то все норм, если нет, то нахожу различия в row, нахожу тем самым суффикс.
	//  Возвращаю букву алфавита + суффикс

	return ""
}
