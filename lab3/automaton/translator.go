package automaton

import "fmt"

type State string
type StateTransitions map[byte][]State

type Machine struct {
	StartState  State
	Transitions map[State]StateTransitions
	FinalStates map[State]struct{}
	Alphabet    map[string]struct{}
}

func New() *Machine {
	return &Machine{
		StartState:  "",
		Transitions: make(map[State]StateTransitions),
		FinalStates: make(map[State]struct{}),
		Alphabet:    make(map[string]struct{}),
	}
}

// Translate TODO: сделать перевод в автомат как в лекции
func (m *Machine) Translate(suffix, prefix map[string]struct{}, table map[string]bool) {
	fmt.Println(suffix, prefix, table)
}

// Membership TODO: доделать функцию
// TODO: беру слово и прохожусь по состояниям - остановился в финальном - принадлежит языку
func (m *Machine) Membership(word string) bool {
	fmt.Println("word:", word)
	return true
}
