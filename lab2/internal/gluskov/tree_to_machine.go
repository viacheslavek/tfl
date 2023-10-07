package gluskov

import (
	"regexp/syntax"
)

type Machine struct {
}

func Translate(st *syntax.Regexp) *Machine {
	return &Machine{}
}
