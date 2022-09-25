package queue

import (
	"errors"
	"fmt"

	"framagit.org/attaboy/go-imm/stack"
)

type Queue[T any] interface {
	Enqueue(T) (Queue[T], error)
	Dequeue() (Queue[T], error)
	Peek() (T, error)
	IsEmpty() bool
	Length() uint
}

type EmptyQueue[T any] struct { }

func NewQueue[T any]() Queue[T] {
	return EmptyQueue[T]{}
}

func (q EmptyQueue[T]) Enqueue(t T) (Queue[T], error) {
	return NonEmptyQueue[T]{
		forwards:  stack.NewStack[T]().Push(t),
		backwards: stack.NewStack[T](),
	}, nil
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

func (q EmptyQueue[T]) Length() uint {
	return 0
}

func (q EmptyQueue[T]) String() string {
	return "fwd: nil\nbwd: nil\n"
}

type NonEmptyQueue[T any] struct {
	forwards, backwards stack.Stack[T]
}

func (q NonEmptyQueue[T]) Enqueue(t T) (Queue[T], error) {
		return NonEmptyQueue[T]{
			forwards:  q.forwards,
			backwards: q.backwards.Push(t),
		}, nil
}

func (q NonEmptyQueue[T]) Dequeue() (Queue[T], error) {
	f, err := q.forwards.Pop()

	if !f.IsEmpty() {
		return NonEmptyQueue[T]{
			forwards:  f,
			backwards: q.backwards,
		}, err
	}

	if !q.backwards.IsEmpty() {
		return NonEmptyQueue[T]{
			forwards:  q.backwards.Reverse(),
			backwards: stack.EmptyStack[T]{},
		}, err
	}

	return EmptyQueue[T]{}, err
}

func (q NonEmptyQueue[T]) Peek() (T, error) {
	return q.forwards.Top()
}

func (q NonEmptyQueue[T]) IsEmpty() bool {
	return false
}

func (q NonEmptyQueue[T]) Length() uint {
	return q.forwards.Depth() + q.backwards.Depth()
}

func (q NonEmptyQueue[T]) String() string {
	return fmt.Sprintf("fwd: %v\nbwd: %v", q.forwards, q.backwards)
}
