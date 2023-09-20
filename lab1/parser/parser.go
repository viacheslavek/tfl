package parser

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/stack"
	"regexp"
	"strings"
)

type Expression struct {
	EPs                       []ExpressionPair
	NameConstructorToConstant map[string]Constructor
	Variables                 map[string]struct{}
}

type ExpressionPair struct {
	Left  string
	Right string
}

type Constructor struct {
	Dimensionality int
	Constants      []string
}

func InitExpression() *Expression {
	return &Expression{
		EPs:                       make([]ExpressionPair, 0),
		NameConstructorToConstant: make(map[string]Constructor),
		Variables:                 make(map[string]struct{}),
	}
}

func (e *Expression) ExtractPair(input string) error {

	inputPairs := strings.Split(input, "\n")

	re := regexp.MustCompile(`([^->]+)->([^->]+)`)

	for _, ip := range inputPairs {
		match := re.FindStringSubmatch(ip)
		if len(match) == 3 {
			left := strings.TrimSpace(match[1])
			right := strings.TrimSpace(match[2])
			pair := ExpressionPair{Left: left, Right: right}
			e.EPs = append(e.EPs, pair)
		} else {
			return fmt.Errorf(basePairFail + "\n" + arrowError)
		}
	}

	return nil
}

func (e *Expression) ParseExpressionsToLinearRepresentation() ([]ExpressionPair, error) {

	linearPair := make([]ExpressionPair, len(e.EPs))
	for i, p := range e.EPs {
		var err error
		linearPair[i].Left, err = e.parseOneFunctionToLinearRepresentation(p.Left)
		linearPair[i].Right, err = e.parseOneFunctionToLinearRepresentation(p.Right)
		if err != nil {
			return make([]ExpressionPair, 0), fmt.Errorf(parseError, err)
		}
	}

	return linearPair, nil
}

func (e *Expression) parseOneFunctionToLinearRepresentation(expr string) (string, error) {

	re := regexp.MustCompile(`[(),]|\w+`)

	// Разбил отдельно на имена конструкторов и переменных, скобки и запятые
	parts := re.FindAllString(expr, -1)

	stackExpr := stack.InitStackString()

	for _, p := range parts {
		switch p {
		case "(":
			if err := e.openBracketCase(stackExpr, p); err != nil {
				return "", err
			}
		case ")":
			if err := e.closeBracketCase(stackExpr); err != nil {
				return "", err
			}
		case ",":
			continue
		default:
			stackExpr.Push(p)
		}
	}

	if stackExpr.Size() != 1 {
		return "", fmt.Errorf("stack must contain one element, it contain %d", stackExpr.Size())
	}

	return stackExpr.Pop()
}

func (e *Expression) openBracketCase(s *stack.Stack[string], p string) error {
	_, err := s.Back()
	if err != nil {
		return fmt.Errorf("in case '(' with element %s error %+v\n", p, err)
	}
	s.Push(p)

	return nil
}

func (e *Expression) closeBracketCase(s *stack.Stack[string]) error {

	curVariables := make([]string, 0)

	for countElem := 0; countElem < 3; countElem++ {
		curElem, err := s.Pop()
		if err != nil {
			return fmt.Errorf("in case ')' in loop iterators %d element has error %+v\n", countElem, err)
		}
		if curElem == "(" {
			constructor, cErr := s.Pop()
			if cErr != nil {
				return fmt.Errorf("in case ')' in pop constructor name has error %+v\n", cErr)
			}

			form, fErr := e.composeLinearForm(constructor, curVariables)
			if fErr != nil {
				return fmt.Errorf("in case ')' in compose form has error %+v\n", fErr)
			}

			s.Push(form)
			return nil
		}
		curVariables = append(curVariables, curElem)
	}

	return fmt.Errorf("in case ')' constructor has more than 2 elements\n")
}

func (e *Expression) composeLinearForm(constructor string, curVariables []string) (string, error) {

	fmt.Println("COMPOSE FORM")

	// работа с добавлением реальных переменных
	// использую грязный хак: у переменной по условию нет впереди себя открывающей скобки
	// если открывающая скобка есть - то это уже линейное выражение
	// соответственно проверяю на наличие скобки и кладу их в set (в Golang это мапа пустых структур)

	for _, cv := range curVariables {
		if len(cv) == 0 {
			return "", fmt.Errorf("one of variables is an empty")
		}
		if cv[0] != '(' {
			e.Variables[cv] = struct{}{}
		}
	}

	fmt.Println("compose Linear form", constructor, curVariables)

	// если конструктор уже лежал в мапе, то я сравниваю размерности
	// если они совпадают, то константы уже заданы, иначе - создаю список констант и кладу их в мапу
	if _, ok := e.NameConstructorToConstant[constructor]; ok {
		if e.NameConstructorToConstant[constructor].Dimensionality != len(curVariables) {
			return "", fmt.Errorf(
				"dimensionality constructors '%s' isn`t equal. was: %d, given: %d",
				constructor, e.NameConstructorToConstant[constructor].Dimensionality, len(curVariables))
		}
	} else {
		e.NameConstructorToConstant[constructor] = Constructor{
			Dimensionality: len(curVariables),
			Constants:      generateConstants(constructor, len(curVariables)+1),
		}
	}

	fmt.Println("constant", e.NameConstructorToConstant[constructor])

	return getLinearForm(e.NameConstructorToConstant[constructor], curVariables), nil
}

func generateConstants(constructName string, countVar int) []string {
	constants := make([]string, countVar)
	for i := 0; i < countVar; i++ {
		constants[i] = fmt.Sprintf("%s_%d", constructName, i)
	}
	return constants
}

func getLinearForm(c Constructor, variable []string) string {

	var linearForm string

	switch c.Dimensionality {
	case 0:
		// const_0
		linearForm = fmt.Sprintf("(%s)", c.Constants[0])
	case 1:
		// x * const_1 + const_0
		linearForm = fmt.Sprintf("(%s * %s + %s)", variable[0], c.Constants[1], c.Constants[0])
	case 2:
		// y * const_2 + x * const_1 + const_0
		linearForm = fmt.Sprintf(
			"(%s * %s + %s * %s + %s)", variable[1], c.Constants[2], variable[0], c.Constants[1], c.Constants[0])
	}

	return linearForm
}
