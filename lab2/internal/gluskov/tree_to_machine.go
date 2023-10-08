package gluskov

import (
	"regexp/syntax"
)

type State int
type StateTransitions map[rune][]State

type Machine struct {
	StartState   int
	FinalStates  []int
	Transitions  map[State]StateTransitions
	StateCounter int
}

func Translate(st *syntax.Regexp) *Machine {
	machine := &Machine{
		StartState:   0,
		FinalStates:  []int{},
		Transitions:  make(map[State]StateTransitions),
		StateCounter: 1,
	}

	machine.buildMachine(st, machine.StartState)

	return machine
}

func (m *Machine) buildMachine(node *syntax.Regexp, currentState int) {
	switch node.Op {
	case syntax.OpLiteral:
		m.addLiteral(currentState)
	case syntax.OpConcat:
		m.addConcat(currentState)
	case syntax.OpAlternate:
		m.addAlternate(currentState)
	case syntax.OpStar:
		m.addStar(currentState)
	case syntax.OpCapture:
		m.addCapture(currentState)
	case syntax.OpCharClass:
		m.addCharClass(currentState)
	}
}

func (m *Machine) addTransition(fromState, toState State, symbol rune) {
	if _, exists := m.Transitions[fromState]; !exists {
		m.Transitions[fromState] = make(StateTransitions)
	}
	m.Transitions[fromState][symbol] = append(m.Transitions[fromState][symbol], toState)
}

func (m *Machine) addState() State {
	newState := State(m.StateCounter)
	m.StateCounter++
	return newState
}

// TODO: сделать корректное добавление буквы - новое состояние и соединяем
// TODO: подумать больше
func (m *Machine) addLiteral(currentState int) {

}

// TODO: сделать корректное добавление конкатенации - просто соединяем состояния
// TODO: подумать больше
func (m *Machine) addConcat(currentState int) {

}

// TODO: сделать корректное добавление альтернативы - еще раз в тетрадь
// TODO: подумать больше
func (m *Machine) addAlternate(currentState int) {

}

// TODO: сделать корректное добавление звезды клини - еще раз нарисовать в тетради
// TODO: подумать больше
func (m *Machine) addStar(currentState int) {

}

// TODO: сделать корректное добавление захвата - просто провалиться вниз
// TODO: подумать больше
func (m *Machine) addCapture(currentState int) {

}

// TODO: сделать корректное добавление самой внутренней альтернативы (это если (a|b))
// TODO: подумать больше
func (m *Machine) addCharClass(currentState int) {

}
