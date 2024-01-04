package tables

type Angluin struct {
	alphabet     []byte
	suffix       map[string]struct{}
	prefix       map[string]struct{}
	extendPrefix map[string]struct{}
	table        map[string]bool
	extendTable  map[string]bool
}

// New TODO: дописать это до нормы
func New() *Angluin {
	return &Angluin{}
}

// Run TODO: Это делаю уже в последнюю очередь:
// Запускаю прогон, если таблица констистента и полна, то кидаю в учителя
// если все норм, то отдаю автомат, если нет, то дальше
func (a *Angluin) Run() {

}

// TODO: проверяю констистентость как в презе
func (a *Angluin) consistent() {

}

// TODO: проверяю полноту как в презе
func (a *Angluin) closed() {

}
