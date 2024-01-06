package mat

import "github.com/VyacheslavIsWorkingNow/tfl/lab3/oracle"

type MAT struct {
	oracle     oracle.Oracle
	countWords int
}

func New(o oracle.Oracle, countWords int) *MAT {
	return &MAT{
		oracle:     o,
		countWords: countWords,
	}
}

func generateAlphabetPermutations(countWords int) []string {
	return []string{}
}

func Equal() (bool, string) {
	return true, ""
}
