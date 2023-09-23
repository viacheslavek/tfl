package parser

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/stack"
)

func convertFromInfixToPrefixNotation(elems []string) (string, error) {

	s := stack.InitStackString()

	for i := 0; i < len(elems); i++ {
		if elems[i] == "+" || elems[i] == "*" {
			prevElem, err := s.Pop()
			if err != nil {
				return "", fmt.Errorf("can`t get previous var %w", err)
			}
			newVar := fmt.Sprintf("(%s %s %s)", elems[i], prevElem, elems[i+1])
			s.Push(newVar)
			i += 1
		} else {
			s.Push(elems[i])
		}
	}

	if s.Size() != 1 {
		return "", fmt.Errorf("stack is not contained one element")
	}

	return s.Pop()
}

// Коэффициенты слева перед переменной >= коэффициенты справа перед переменной для всех переменных
func (e *Expression) makeFirstInequality() {

}

// Коэффициенты слева свободных членов >= коэффициенты справа свободных членов
func (e *Expression) makeSecondInequality() {

}

// Коэффициенты слева всего выражения через or > коэффициенты справа всего выражения через or
func (e *Expression) makeThirdInequality() {

}

// Коэффициенты конструкторов через and для переменных >= 1 и для констант >= 0
func (e *Expression) makeForthInequality() {

}

// Коэффициенты конструкторов через and, в котором внутри еще or (переменных > 1 и для констант > 0)
func (e *Expression) makeFifthInequality() {

}
