package stack

import (
	"errors"
)

// Stack is an immutable implementation of a generic stack.
type Stack[T any] interface {
	// Push adds an element t to the top of the stack.
	// This method is pure, the return value must be used.
	Push(t T) Stack[T]

	// Pop removes an element from the stack.
	// If the stack is empty, a non-nil error is returned.
	// This method is pure, the return value must be used.
	Pop() (Stack[T], error)

	// Top returns the top element of the stack.
	// If the stack is empty, err is non-nil.
	Top() (T, error)

	// Depth returns the depth of the stack.
	Depth() uint
}

// EmptyStack is an empty stack.
type EmptyStack[T any] struct{}

// NewStack creates an empty stack.
func NewStack[T any]() Stack[T] {
	return EmptyStack[T]{}
}

func (e EmptyStack[T]) Push(t T) Stack[T] {
	return NonEmptyStack[T]{
		topEl: t,
		tail:  e,
	}
}

func (e EmptyStack[T]) Pop() (Stack[T], error) {
	return e, errors.New("cannot pop empty stack")
}

func (e EmptyStack[T]) Top() (T, error) {
	var noop T
	return noop, errors.New("cannot top on empty stack")
}

func (e EmptyStack[T]) Depth() uint {
	return 0
}

// NonEmptyStack is a non-empty stack.
type NonEmptyStack[T any] struct {
	topEl T
	tail  Stack[T]
}

// NewNonEmptyStack creates a non-empty stack, using top as the top and
// previous as the stack's tail.
func NewNonEmptyStack[T any](top T, previous Stack[T]) Stack[T] {
	return NonEmptyStack[T]{
		topEl: top,
		tail:  previous,
	}
}

func (e NonEmptyStack[T]) Push(t T) Stack[T] {
	return NonEmptyStack[T]{
		topEl: t,
		tail:  e,
	}
}

func (e NonEmptyStack[T]) Pop() (Stack[T], error) {
	return e.tail, nil
}

func (e NonEmptyStack[T]) Top() (T, error) {
	return e.topEl, nil
}

func (e NonEmptyStack[T]) Depth() uint {
	return 1 + e.tail.Depth()
}
