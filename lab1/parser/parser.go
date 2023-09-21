package parser

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/tfl/lab1/stack"
	"log"
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

func Parse(input string) error {
	expr := InitExpression()

	if err := expr.ExtractPair(input); err != nil {
		return fmt.Errorf(parsePairError, err)
	}

	if err := expr.ParseExpressionsToLinearRepresentation(); err != nil {
		return fmt.Errorf(parseRepresentationError, err)
	}

	panic("я еще не готова.")

	return nil
}

func (e *Expression) ExtractPair(input string) error {

	log.Printf("start parsing\n")

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

	log.Printf("after spliting into pairs:\n%s", e.ToStringLinearExpression())

	return nil
}

func (e *Expression) ToStringLinearExpression() string {
	linear := ""
	for i, expr := range e.EPs {
		linear += fmt.Sprintf("%d:\nleft: %s\nright: %s\n", i, expr.Left, expr.Right)
	}
	linear += "\n"
	return linear
}

func (e *Expression) ToStringConstructions() string {
	constr := ""
	for name, constant := range e.NameConstructorToConstant {
		constr += fmt.Sprintf("name: %s, constants: %+v\n", name, constant)
	}
	constr += "\n"
	return constr
}

func (e *Expression) ToStringVariables() string {
	variables := ""
	for v := range e.Variables {
		variables += fmt.Sprintf(", %s", v)
	}
	variables += "\n"
	return variables[2:]
}

func (e *Expression) ParseExpressionsToLinearRepresentation() error {

	linearPair := make([]ExpressionPair, len(e.EPs))
	for i, p := range e.EPs {
		var err error
		linearPair[i].Left, err = e.parseOneFunctionToLinearRepresentation(p.Left)
		linearPair[i].Right, err = e.parseOneFunctionToLinearRepresentation(p.Right)
		if err != nil {
			return fmt.Errorf(parseError, err)
		}
	}

	e.EPs = linearPair

	log.Printf("after parsing into linear expression\n")
	log.Printf("linear expression: \n%s", e.ToStringLinearExpression())
	log.Printf("constructor and constants:\n%s", e.ToStringConstructions())
	log.Printf("variables:\n%s", e.ToStringVariables())

	return nil
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
		return "", fmt.Errorf(stackSizeError, stackExpr.Size())
	}

	return stackExpr.Pop()
}

func (e *Expression) openBracketCase(s *stack.Stack[string], p string) error {
	_, err := s.Back()
	if err != nil {
		return fmt.Errorf(openBracketError, p, err)
	}
	s.Push(p)

	return nil
}

func (e *Expression) closeBracketCase(s *stack.Stack[string]) error {

	curVariables := make([]string, 0)

	for countElem := 0; countElem < 3; countElem++ {
		curElem, err := s.Pop()
		if err != nil {
			return fmt.Errorf(closeBracketStackLoopError, countElem, err)
		}
		if curElem == "(" {
			constructor, cErr := s.Pop()
			if cErr != nil {
				return fmt.Errorf(closeBracketStackConstrError, cErr)
			}

			form, fErr := e.composeLinearForm(constructor, curVariables)
			if fErr != nil {
				return fmt.Errorf(closeBracketComposeError, fErr)
			}

			s.Push(form)
			return nil
		}
		curVariables = append(curVariables, curElem)
	}

	return fmt.Errorf(constructorTooMuchParams)
}

func (e *Expression) composeLinearForm(constructor string, curVariables []string) (string, error) {

	// работа с добавлением реальных переменных
	// использую грязный хак: у переменной по условию нет впереди себя открывающей скобки
	// если открывающая скобка есть - то это уже линейное выражение
	// соответственно проверяю на наличие скобки и кладу их в set (в Golang это мапа пустых структур)

	for _, cv := range curVariables {
		if len(cv) == 0 {
			return "", fmt.Errorf(emptyVariable)
		}
		if cv[0] != '(' {
			e.Variables[cv] = struct{}{}
		}
	}

	// если конструктор уже лежал в мапе, то я сравниваю размерности
	// если они совпадают, то константы уже заданы, иначе - создаю список констант и кладу их в мапу
	if _, ok := e.NameConstructorToConstant[constructor]; ok {
		if e.NameConstructorToConstant[constructor].Dimensionality != len(curVariables) {
			return "",
				fmt.Errorf(dimensionalNotEqual,
					constructor, e.NameConstructorToConstant[constructor].Dimensionality, len(curVariables),
				)
		}
	} else {
		e.NameConstructorToConstant[constructor] = Constructor{
			Dimensionality: len(curVariables),
			Constants:      generateConstants(constructor, len(curVariables)+1),
		}
	}

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

func (e *Expression) BringingSuchForLinearForms() ([]ExpressionPair, error) {

	bringingPair := make([]ExpressionPair, len(e.EPs))
	for i, p := range e.EPs {
		var err error
		bringingPair[i].Left, err = e.BringingLinearForm(p.Left)
		bringingPair[i].Right, err = e.BringingLinearForm(p.Right)
		if err != nil {
			return make([]ExpressionPair, 0), fmt.Errorf(parseError, err)
		}
	}

	return bringingPair, nil
}

func (e *Expression) BringingLinearForm(expr string) (string, error) {

	fmt.Println("expr:", expr)

	// TODO: раскрываем скобки умножением

	linearFormWithoutBrackets, wbErr := e.openBracketsMultiplication(expr)
	if wbErr != nil {
		return "", fmt.Errorf("can't open multiplicative brackets, %w", wbErr)
	}

	fmt.Println("after open brackets", linearFormWithoutBrackets)

	// TODO: приводим подобные

	return "", nil
}

func (e *Expression) openBracketsMultiplication(expr string) (string, error) {

	// Убираем скобки слева и справа от выражения
	expr = expr[1 : len(expr)-1]
	fmt.Println(expr)

	re := regexp.MustCompile(`[()*+]|\w+`)

	// Разбил отдельно на имена переменных, константы, скобки, знаки сложения и умножения
	parts := re.FindAllString(expr, -1)

	for _, p := range parts {
		fmt.Printf("%s ", p)
	}
	fmt.Println()

	stackExpr := stack.InitStackString()

	for i := 0; i < len(parts); i++ {
		switch parts[i] {
		case ")":
			// за закрывающей скобкой всегда следует знак умножения и то, на что мы умножаем
			if err := e.openDistributivity(stackExpr, parts[i+2]); err != nil {
				return "", fmt.Errorf("in case ')' was error: %w", err)
			}
			i += 2
		default:
			stackExpr.Push(parts[i])
		}
	}

	return "", nil
}

func (e *Expression) openDistributivity(s *stack.Stack[string], multiplier string) (err error) {

	var elem string
	terms := make([]string, 0)
	operations := make([]string, 0)

	for elem != "(" {
		elem, err = s.Pop()
		if err != nil {
			return fmt.Errorf("in loop pop was %w", err)
		}
		if elem == "+" || elem == "*" {
			operations = append(operations, elem)
		} else if elem == "(" {
			continue
		} else {
			terms = append(terms, elem)
		}
	}

	s.Push(constructDistributivity(terms, operations, multiplier))

	return nil
}

func constructDistributivity(terms, operations []string, multiplier string) string {
	// нужно изобрести велосипед с корректным перемножением
	return ""
}
