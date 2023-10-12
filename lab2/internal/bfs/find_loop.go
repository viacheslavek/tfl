package bfs

import "github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/gluskov"

type dfsParam struct {
	machine gluskov.Machine
	cycles  [][]gluskov.State
	visited map[gluskov.State]bool
	path    []gluskov.State
}

func FindCycles(machine gluskov.Machine) [][]gluskov.State {

	dfsp := newDfsParam(machine)

	for i := 0; i < machine.StateCounter; i++ {
		if !dfsp.visited[gluskov.State(i)] {
			dfsp.dfs(gluskov.State(i))
		}
	}

	return dfsp.cycles
}

func newDfsParam(machine gluskov.Machine) *dfsParam {
	return &dfsParam{
		machine: machine,
		cycles:  make([][]gluskov.State, 0),
		visited: make(map[gluskov.State]bool),
		path:    make([]gluskov.State, 0),
	}
}

func (dp *dfsParam) dfs(currentState gluskov.State) {
	if dp.visited[currentState] {
		for i, state := range dp.path {
			if state == currentState {
				cycle := dp.path[i:]
				dp.cycles = append(dp.cycles, cycle)
				return
			}
		}
		return
	}

	dp.visited[currentState] = true
	dp.path = append(dp.path, currentState)

	transitions := dp.machine.Transitions[currentState]
	for _, nextStates := range transitions {
		for _, nextState := range nextStates {
			dp.dfs(nextState)
		}
	}

	dp.path = dp.path[:len(dp.path)-1]
}
