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
func (e *Expression) MakeFirstInequality() ([]string, error) {

	firstInequalities := make([]string, 0)

	for _, pairs := range e.EPs {
		inequality, err := getFirst(pairs)
		if err != nil {
			return make([]string, 0), fmt.Errorf("can't convert first expr ineq %w", err)
		}
		firstInequalities = append(firstInequalities, inequality...)
	}

	return firstInequalities, nil
}

// TODO : Эта функция явно портит весь мой код - отрефакторить

func getFirst(ep ExpressionPair) ([]string, error) {

	variablesStr := make([]string, 0)

	usages := make(map[string]struct{})

	usages[" "] = struct{}{}

	for key, value := range ep.Left.cAv {
		if _, ok := usages[key]; ok {
			continue
		}
		usages[key] = struct{}{}
		valueRight, ok := ep.Right.cAv[key]
		if ok != true {
			valueRight = []string{"0"}
		}
		varStr, err := getFirstExpression(value, valueRight)
		if err != nil {
			return make([]string, 0), fmt.Errorf("can't convert first expr %w", err)
		}
		variablesStr = append(variablesStr, varStr)
	}

	for key, value := range ep.Right.cAv {
		if _, ok := usages[key]; ok {
			continue
		}
		usages[key] = struct{}{}
		valueLeft, ok := ep.Left.cAv[key]
		if ok != true {
			valueLeft = []string{"0"}
		}
		varStr, err := getFirstExpression(valueLeft, value)
		if err != nil {
			return make([]string, 0), fmt.Errorf("can't convert first expr %w", err)
		}
		variablesStr = append(variablesStr, varStr)
	}

	return variablesStr, nil
}

func getFirstExpression(valueL, valueR []string) (string, error) {

	left, lErr := convertFromInfixToPrefixNotation(valueL)
	if lErr != nil {
		return "", fmt.Errorf("can't convert left value to prefix %w", lErr)
	}

	right, rErr := convertFromInfixToPrefixNotation(valueR)
	if lErr != nil {
		return "", fmt.Errorf("can't convert right value to prefix %w", rErr)
	}

	return fmt.Sprintf(
		"(assert (>= %s %s))",
		left,
		right,
	), nil
}

// Коэффициенты слева свободных членов >= коэффициенты справа свободных членов
func (e *Expression) MakeSecondInequality() ([]string, error) {
	secondInequalities := make([]string, 0)

	for _, pairs := range e.EPs {
		inequality, err := getSecond(pairs)
		if err != nil {
			return make([]string, 0), fmt.Errorf("can't convert second expr ineq %w", err)
		}
		secondInequalities = append(secondInequalities, inequality)
	}

	return secondInequalities, nil
}

func getSecond(ep ExpressionPair) (string, error) {

	valueFreedomLeft, okL := ep.Left.cAv[" "]
	if okL != true {
		valueFreedomLeft = []string{"0"}
	}

	valueFreedomRight, okR := ep.Right.cAv[" "]
	if okR != true {
		valueFreedomRight = []string{"0"}
	}

	return getSecondExpression(valueFreedomLeft, valueFreedomRight)
}

func getSecondExpression(valueL, valueR []string) (string, error) {

	left, lErr := convertFromInfixToPrefixNotation(valueL)
	if lErr != nil {
		return "", fmt.Errorf("can't convert left value to prefix %w", lErr)
	}

	right, rErr := convertFromInfixToPrefixNotation(valueR)
	if lErr != nil {
		return "", fmt.Errorf("can't convert right value to prefix %w", rErr)
	}

	return fmt.Sprintf(
		"(assert (>= %s %s))",
		left,
		right,
	), nil
}

// Коэффициенты слева всего выражения через or > коэффициенты справа всего выражения через or
func (e *Expression) MakeThirdInequality() ([]string, error) {
	thirdInequalities := make([]string, 0)

	for _, pairs := range e.EPs {

		inequality, err := getThird(pairs)
		if err != nil {
			return make([]string, 0), fmt.Errorf("can't convert first expr ineq %w", err)
		}
		thirdInequalities = append(thirdInequalities, inequality)
	}

	return thirdInequalities, nil
}

// слишком много дублирования с первой - потом поправить

// TODO : Эта функция явно портит весь мой код - отрефакторить хахахаха - уже второй раз
// вот так и сдавай за два часа до дедлайна

func getThird(ep ExpressionPair) (string, error) {

	variablesStr := make([]string, 0)

	usages := make(map[string]struct{})

	for key, value := range ep.Left.cAv {
		if _, ok := usages[key]; ok {
			continue
		}
		usages[key] = struct{}{}
		valueRight, ok := ep.Right.cAv[key]
		if ok != true {
			valueRight = []string{"0"}
		}
		varStr, err := getThirdExpression(value, valueRight)
		if err != nil {
			return "", fmt.Errorf("can't convert third expr %w", err)
		}
		variablesStr = append(variablesStr, varStr)
	}

	for key, value := range ep.Right.cAv {
		if _, ok := usages[key]; ok {
			continue
		}
		usages[key] = struct{}{}
		valueLeft, ok := ep.Left.cAv[key]
		if ok != true {
			valueLeft = []string{"0"}
		}
		varStr, err := getThirdExpression(valueLeft, value)
		if err != nil {
			return "", fmt.Errorf("can't convert third expr %w", err)
		}
		variablesStr = append(variablesStr, varStr)
	}

	// TODO: низкую логику в отдельную функцию

	thirdExpr := "(assert (or"
	for _, p := range variablesStr {
		thirdExpr += " "
		thirdExpr += p
	}
	thirdExpr += "))"

	return thirdExpr, nil
}

func getThirdExpression(valueL, valueR []string) (string, error) {

	left, lErr := convertFromInfixToPrefixNotation(valueL)
	if lErr != nil {
		return "", fmt.Errorf("can't convert left value to prefix %w", lErr)
	}

	right, rErr := convertFromInfixToPrefixNotation(valueR)
	if lErr != nil {
		return "", fmt.Errorf("can't convert right value to prefix %w", rErr)
	}

	return fmt.Sprintf(
		"(> %s %s)",
		left,
		right,
	), nil
}

// Коэффициенты конструкторов через and для переменных >= 1 и для констант >= 0
func (e *Expression) MakeForthInequality() string {

	forthInequalitiesPred := make([]string, 0)
	for _, nmt := range e.NameConstructorToConstant {
		forthInequalitiesPred = append(forthInequalitiesPred, getForth(nmt))
	}

	exp := "(assert (and"

	for _, fip := range forthInequalitiesPred {
		exp += " "
		exp += fip
	}
	exp += "))"

	return exp
}

func getForth(constructor Constructor) string {

	switch constructor.Dimensionality {
	case 0:
		return fmt.Sprintf("(>= %s 0)", constructor.Constants[0])
	case 1:
		return fmt.Sprintf("(>= %s 0) (>= %s 1)", constructor.Constants[0], constructor.Constants[1])
	case 2:
		return fmt.Sprintf("(>= %s 0) (>= %s 1) (>= %s 1)",
			constructor.Constants[0], constructor.Constants[1], constructor.Constants[2])
	}
	return ""
}

// Коэффициенты конструкторов через and, в котором внутри еще or (переменных > 1 и для констант > 0)
func (e *Expression) MakeFifthInequality() string {
	fifthInequalitiesPred := make([]string, 0)
	for _, nmt := range e.NameConstructorToConstant {
		fifthInequalitiesPred = append(fifthInequalitiesPred, getFifth(nmt))
	}

	exp := "(assert (and"

	for _, fip := range fifthInequalitiesPred {
		exp += " "
		exp += fip
	}
	exp += "))"

	return exp
}

func getFifth(constructor Constructor) string {

	switch constructor.Dimensionality {
	case 0:
		return fmt.Sprintf("(> %s 0)", constructor.Constants[0])
	case 1:
		return fmt.Sprintf("(or (> %s 0) (> %s 1))", constructor.Constants[0], constructor.Constants[1])
	case 2:
		return fmt.Sprintf("(or (> %s 0) (and (> %s 1) (> %s 1)))",
			constructor.Constants[0], constructor.Constants[1], constructor.Constants[2])
	}
	return ""
}
