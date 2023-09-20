package stack

import "errors"

type Stack[T comparable] struct {
	buffer []T
}

func InitStackString() *Stack[string] {
	return &Stack[string]{
		buffer: make([]string, 0),
	}
}

func (obj *Stack[T]) Push(elem T) {
	obj.buffer = append(obj.buffer, elem)
}

func (obj *Stack[T]) Pop() (T, error) {
	if len(obj.buffer) > 0 {
		elem := obj.buffer[len(obj.buffer)-1]
		obj.buffer = obj.buffer[:len(obj.buffer)-1]
		return elem, nil
	}
	var temp T
	return temp, errors.New("empty buffer")
}

func (obj *Stack[T]) Back() (T, error) {
	if len(obj.buffer) > 0 {
		elem := obj.buffer[len(obj.buffer)-1]
		return elem, nil
	}
	var temp T
	return temp, errors.New("empty buffer")
}

func (obj *Stack[T]) Size() int {
	return len(obj.buffer)
}

func (obj *Stack[T]) Clear() {
	obj.buffer = make([]T, 0)
}
