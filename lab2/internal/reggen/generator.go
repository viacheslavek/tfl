package reggen

import "fmt"

const (
	maxAlphabetNumber = 52
)

type Regexes struct {
	rgp        RegexGeneratorParams
	countRegex int
}

type RegexGeneratorParams struct {
	alphabetSize int
	starHeight   int
	letterCount  int
}

func New(
	alphabetSize int,
	starHeight int,
	letterCount int,
) *RegexGeneratorParams {
	return &RegexGeneratorParams{
		alphabetSize: alphabetSize,
		starHeight:   starHeight,
		letterCount:  letterCount,
	}
}

// TODO: все переделать - генерирую обычное регулярное выражение в виде строки

func (r *Regexes) Generate() []string {

	regexes := make([]string, r.countRegex)

	for i := 0; i < r.countRegex; i++ {
		regexes[i] = r.rgp.Generate()
	}

	return regexes
}

type generateRegexPresentation struct {
	rgp                RegexGeneratorParams
	nestingParentheses int
	nestingStars       int
}

func (rgp *RegexGeneratorParams) Generate() string {

	// This is coping
	grp := generateRegexPresentation{
		rgp: *rgp,
	}

	regex := ""

	grp.recursiveGenerate(&regex)

	return regex
}

func (grp *generateRegexPresentation) recursiveGenerate(regex *string) {
	if grp.rgp.letterCount == 0 {
		return
	}

	grp.rgp.letterCount--

	// TODO: тут создаю рекурсивную функцию
	// где я пока не получил 0 в длине - вызываю
	// но мне надо уметь закрывать скобки
	// то есть надо еще понимать, смогу ли я вообще при заданных параметрах что-то сделать
	// звездная высота требует ((a*)*)* скобок
	// то есть надо сделать внутреннюю структуру, с информацией о вложенности внутри скобочной последовательности
	// и о вложенности внутри звездной высоты

	grp.recursiveGenerate(regex)
}

func getLetter(numLetter int) (string, error) {
	if numLetter > maxAlphabetNumber {
		return "", fmt.Errorf("letter num is too big: max is %d, get: %d", maxAlphabetNumber, numLetter)
	}
	return string('a' - 1 + byte(numLetter)), nil
}
