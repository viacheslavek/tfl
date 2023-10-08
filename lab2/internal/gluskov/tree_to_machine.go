package gluskov

import (
	"fmt"
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

	machine.buildMachine(st, State(machine.StartState))

	return machine
}

func (m *Machine) buildMachine(node *syntax.Regexp, currentState State) State {
	switch node.Op {
	case syntax.OpLiteral:
		return m.addLiteral(currentState, node)
	case syntax.OpConcat:
		return m.addConcat(currentState, node)
	case syntax.OpAlternate:
		return m.addAlternate(currentState, node)
	case syntax.OpStar:
		return m.addStar(currentState, node)
	case syntax.OpCapture:
		return m.addCapture(currentState, node)
	case syntax.OpCharClass:
		return m.addCharClass(currentState, node)
	}
	fmt.Println("вышли за case")
	return currentState
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

func (m *Machine) addLiteral(currentState State, node *syntax.Regexp) State {
	for _, symbol := range node.Rune {
		nextState := m.addState()
		m.addTransition(currentState, nextState, symbol)
		currentState = nextState
	}
	return currentState
}

// TODO: сделать корректное добавление конкатенации - просто соединяем состояния
// TODO: подумать больше
func (m *Machine) addConcat(currentState State, node *syntax.Regexp) State {
	panic("implement me")
}

// TODO: сделать корректное добавление альтернативы - еще раз в тетрадь
// TODO: подумать больше
func (m *Machine) addAlternate(currentState State, node *syntax.Regexp) State {
	panic("implement me")
}

// TODO: сделать корректное добавление звезды клини - еще раз нарисовать в тетради
// TODO: подумать больше
func (m *Machine) addStar(currentState State, node *syntax.Regexp) State {
	panic("implement me")
}

// TODO: сделать корректное добавление захвата - просто провалиться вниз
// TODO: подумать больше
func (m *Machine) addCapture(currentState State, node *syntax.Regexp) State {
	panic("implement me")
}

// TODO: сделать корректное добавление самой внутренней альтернативы (это если (a|b))
// TODO: подумать больше
func (m *Machine) addCharClass(currentState State, node *syntax.Regexp) State {
	panic("implement me")
}
