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

	fmt.Println(a.table)
	a.PrintTable()
	a.addPrefix("a")
	a.updateTableByPrefix("a")
	a.addPrefix("aa")
	fmt.Println(a.table)
	a.PrintTable()
	a.updateTableByPrefix("aa")
	fmt.Println(a.table)
	a.PrintTable()

	a.addSuffix("b")
	a.updateTableBySuffix("b")
	fmt.Println(a.table)
	a.PrintTable()
	a.addSuffix("bb")
	a.updateTableBySuffix("bb")
	fmt.Println(a.table)
	a.PrintTable()

	fmt.Println("end in run")
}

// TODO: проверяю полноту как в презе
func (a *Angluin) closed() {

}

// TODO: проверяю констистентость как в презе
func (a *Angluin) consistent() {

}
