package gluskov

import (
	"fmt"
	"regexp/syntax"
)

type State int
type StateTransitions map[rune][]State

type Machine struct {
	StartState   int
	FinalStates  []State
	Transitions  map[State]StateTransitions
	StateCounter int
}

func BuildMachine(st *syntax.Regexp) *Machine {
	machine := &Machine{
		StartState:   0,
		FinalStates:  make([]State, 0),
		Transitions:  make(map[State]StateTransitions),
		StateCounter: 1,
	}

	machine.handleRegex(st, State(machine.StartState), true)

	return machine
}

func (m *Machine) handleRegex(node *syntax.Regexp, currentState State, isFinal bool) []State {
	switch node.Op {
	case syntax.OpLiteral:
		return m.handleLiteral(currentState, node, isFinal)
	case syntax.OpConcat:
		// TODO: реализовать
		return m.handleConcat(currentState, node, isFinal)
	case syntax.OpAlternate:
		// TODO: реализовать
		states := m.handleAlternate(currentState, node, isFinal)
		fmt.Println(states)
	case syntax.OpStar:
		// TODO: реализовать
		return m.handleStar(currentState, node)
	case syntax.OpCapture:
		return m.handleCapture(currentState, node, isFinal)
	case syntax.OpCharClass:
		// TODO: реализовать
		return m.handleCharClass(currentState, node)
	}
	fmt.Println("вышли за case")
	return []State{currentState}
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

func (m *Machine) addFinal(s State) {
	m.FinalStates = append(m.FinalStates, s)
}

func (m *Machine) handleLiteral(currentState State, node *syntax.Regexp, isFinal bool) []State {
	for _, symbol := range node.Rune {
		nextState := m.addState()
		m.addTransition(currentState, nextState, symbol)
		currentState = nextState
	}
	if isFinal {
		m.addFinal(currentState)
	}
	return []State{currentState}
}

// TODO: левая часть с false, правая часть с true как final
func (m *Machine) handleConcat(currentState State, node *syntax.Regexp, isFinal bool) []State {

	panic("implement me")
}

func (m *Machine) handleAlternate(currentState State, node *syntax.Regexp, isFinal bool) []State {
	fmt.Println("alterNode", node)
	leftState := m.handleRegex(node.Sub[0], currentState, isFinal)
	rightState := m.handleRegex(node.Sub[1], currentState, isFinal)

	if isFinal {
		m.addFinal(leftState[0])
		m.addFinal(rightState[0])
	}
	return []State{leftState[0], rightState[0]}
}

// TODO: сделать корректное добавление звезды клини
func (m *Machine) handleStar(currentState State, node *syntax.Regexp) []State {
	panic("implement me")
}

func (m *Machine) handleCapture(currentState State, node *syntax.Regexp, isFinal bool) []State {
	if len(node.Sub) != 1 {
		panic("Длина node.Sub в захвате не равна 1 -> такой случай я не рассматривал")
	}
	return m.handleRegex(node.Sub[0], currentState, isFinal)
}

// TODO: Здесь все не сложно, как с обычной альтернативой
func (m *Machine) handleCharClass(currentState State, node *syntax.Regexp) []State {
	panic("implement me")
}
