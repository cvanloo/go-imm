package stack

import (
	"errors"
	"fmt"
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

	// IsEmpty returns true if the stack is empty, false if it's non-empty.
	IsEmpty() bool

	// Reverse the stack
	Reverse() Stack[T]
}

// EmptyStack is an empty stack.
type EmptyStack[T any] struct{}

// NewStack creates an empty stack.
func NewStack[T any]() Stack[T] {
	return EmptyStack[T]{}
}

func (s EmptyStack[T]) Push(t T) Stack[T] {
	return NonEmptyStack[T]{
		topEl: t,
		tail:  s,
	}
}

func (s EmptyStack[T]) Pop() (Stack[T], error) {
	return s, errors.New("cannot pop empty stack")
}

func (s EmptyStack[T]) Top() (T, error) {
	var noop T
	return noop, errors.New("cannot top on empty stack")
}

func (s EmptyStack[T]) Depth() uint {
	return 0
}

func (s EmptyStack[T]) IsEmpty() bool {
	return true
}

func (s EmptyStack[T]) Reverse() Stack[T] {
	return s
}

func (s EmptyStack[T]) String() string {
	return "nil"
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

func (s NonEmptyStack[T]) Push(t T) Stack[T] {
	return NonEmptyStack[T]{
		topEl: t,
		tail:  s,
	}
}

func (s NonEmptyStack[T]) Pop() (Stack[T], error) {
	return s.tail, nil
}

func (s NonEmptyStack[T]) Top() (T, error) {
	return s.topEl, nil
}

func (s NonEmptyStack[T]) Depth() uint {
	return 1 + s.tail.Depth()
}

func (s NonEmptyStack[T]) IsEmpty() bool {
	return false
}

func (s NonEmptyStack[T]) Reverse() Stack[T] {
	r := NewStack[T]()
	var cur Stack[T] = s
	for !cur.IsEmpty() {
		t, _ := cur.Top()
		r = r.Push(t)
		cur, _ = cur.Pop()
	}
	return r
}

func (s NonEmptyStack[T]) String() string {
	t, _ := s.Top()
	return fmt.Sprintf("%v, %v", t, s.tail)
}
