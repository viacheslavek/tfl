package mat

import (
	"log"

	"github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"
)

// Генератор получился отличный!

type MAT struct {
	oracle      oracle.Oracle
	maxLenWords int
}

func New(o oracle.Oracle, maxLenWords int) *MAT {
	return &MAT{
		oracle:      o,
		maxLenWords: maxLenWords,
	}
}

func (mat *MAT) GenerateWords() []string {

	words := make([]string, 0)

	for i := 1; i <= mat.maxLenWords; i++ {
		words = append(words, generateCombinations(mat.oracle.GetAlphabet(), i)...)
	}

	log.Printf("Generated %d words\n", len(words))

	return words
}

func generateCombinations(alphabet []byte, length int) []string {
	var resultByte [][]byte
	generateCombination(alphabet, length, []byte{}, &resultByte)

	resultString := make([]string, len(resultByte))

	for i := 0; i < len(resultByte); i++ {
		resultString[i] = string(resultByte[i])
	}

	return resultString
}

func generateCombination(alphabet []byte, remaining int, current []byte, result *[][]byte) {
	if remaining == 0 {
		*result = append(*result, append([]byte{}, current...))
		return
	}

	for _, char := range alphabet {
		generateCombination(alphabet, remaining-1, append(current, char), result)
	}
}

func (mat *MAT) Membership(word string) bool {
	return mat.oracle.BelongLanguage(word)
}

// Equivalence TODO: Когда напишу перевод в автомат, то здесь на вход принимаю автомат, генерирую слова
//
//	и эти слова даю на вход автомату и оракулу, если со всеми словами все хорошо - отлично, если нет - выдаю слово
func (mat *MAT) Equivalence() (bool, string) {
	return true, ""
}

func (mat *MAT) GetAlphabet() []byte {
	return mat.oracle.GetAlphabet()
}
