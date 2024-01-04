package tables

// TODO: Просто в suffix set добавляем новую строчку
func (a *Angluin) addSuffix() {

}

// TODO: В prefix set добавляем новую строчку и все его префиксы, если не существуют
func (a *Angluin) addPrefix() {

}

// TODO: Добавляем новые префиксы в ту таблицу: для этого к новым префиксам добавляем буквы из алфавита.
// Проходимся по новым добавленым префиксам и если есть совпадение с префиксами префикса, то удаляем их из set
func (a *Angluin) addExtendPrefix() {

}

// TODO: Если добавили префиксы, то создаём новые значения. Аналогично с суффиксами. Создание - просто проход
func (a *Angluin) updateTable() {

}

// Если что-то удаляем из префиксов S расширенной, то удаляем и из мапы. Если добавляем суффикс,
// то добавляем новые значения в мапу. Аналогично с новыми значениями префиксов.
func (a *Angluin) updateExtendTable() {

}

// PrintTable TODO: сделать
func (a *Angluin) PrintTable() {

}

// PrintExtendTable TODO: сделать
func (a *Angluin) PrintExtendTable() {

}

// PrintPrefix TODO: сделать
func (a *Angluin) PrintPrefix() {

}

// PrintExtendPrefix TODO: сделать
func (a *Angluin) PrintExtendPrefix() {

}

// PrintSuffix TODO: сделать
func (a *Angluin) PrintSuffix() {

}
