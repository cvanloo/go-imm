package queue

import (
	"errors"
	"fmt"

	"framagit.org/attaboy/go-imm/stack"
)

type Queue[T any] interface {
	Enqueue(T) Queue[T]
	Dequeue() (Queue[T], error)
	Peek() (T, error)
	IsEmpty() bool
	//Length() uint
	//Capacity() uint
}

type EmptyQueue[T any] struct{}

func (q EmptyQueue[T]) Enqueue(t T) Queue[T] {
	return NonFullQueue[T]{
		forwards:  stack.NewStack[T]().Push(t),
		backwards: stack.NewStack[T](),
	}
}

func (q EmptyQueue[T]) Dequeue() (Queue[T], error) {
	return nil, errors.New("cannot dequeue an empty queue")
}

func (q EmptyQueue[T]) Peek() (T, error) {
	var noop T
	return noop, errors.New("cannot peek an empty queue")
}

func (q EmptyQueue[T]) IsEmpty() bool {
	return true
}

func (q EmptyQueue[T]) String() string {
	return "fwd: nil\nbwd: nil\n"
}

type NonFullQueue[T any] struct {
	forwards, backwards stack.Stack[T]
}

func (q NonFullQueue[T]) Enqueue(t T) Queue[T] {
	return NonFullQueue[T]{
		forwards:  q.forwards,
		backwards: q.backwards.Push(t),
	}
}

func (q NonFullQueue[T]) Dequeue() (Queue[T], error) {
	f, err := q.forwards.Pop()

	if !f.IsEmpty() {
		return NonFullQueue[T]{
			forwards:  f,
			backwards: q.backwards,
		}, err
	}

	if q.backwards.IsEmpty() {
		return EmptyQueue[T]{}, nil
	}

	return NonFullQueue[T]{
		forwards:  q.backwards.Reverse(),
		backwards: stack.EmptyStack[T]{},
	}, nil
}

func (q NonFullQueue[T]) Peek() (T, error) {
	return q.forwards.Top()
}

func (q NonFullQueue[T]) IsEmpty() bool {
	return false
}

func (q NonFullQueue[T]) String() string {
	return fmt.Sprintf("fwd: %v\nbwd: %v", q.forwards, q.backwards)
}

type FullQueue[T any] struct{}
