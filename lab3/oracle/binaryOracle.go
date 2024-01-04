package oracle

type BinaryOracle struct {
	path string
}

func NewBinaryOracle(path string) *BinaryOracle {
	return &BinaryOracle{path: path}
}

func (bo BinaryOracle) BelongLanguage(word string) bool {
	return matchedBinary(word, bo.path)
}

func matchedBinary(word, path string) bool {
	// TODO: вызываю бинарный файл со словом word по пути path и смотрю на ответ
	return false
}
