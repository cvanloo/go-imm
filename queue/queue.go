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
	Capacity() uint
}

type EmptyQueue[T any] struct {
	capacity uint
}

func NewQueue[T any](capacity uint) Queue[T] {
	return EmptyQueue[T]{capacity: capacity}
}

func (q EmptyQueue[T]) Enqueue(t T) (Queue[T], error) {
	return NonFullQueue[T]{
		capacity:  q.capacity,
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

func (q EmptyQueue[T]) Capacity() uint {
	return q.capacity
}

func (q EmptyQueue[T]) String() string {
	return "fwd: nil\nbwd: nil\n"
}

type NonFullQueue[T any] struct {
	capacity            uint
	forwards, backwards stack.Stack[T]
}

func (q NonFullQueue[T]) Enqueue(t T) (Queue[T], error) {
	if q.Length() < q.capacity-1 {
		return NonFullQueue[T]{
			capacity:  q.capacity,
			forwards:  q.forwards,
			backwards: q.backwards.Push(t),
		}, nil
	}

	return FullQueue[T]{
		capacity:  q.capacity,
		forwards:  q.forwards,
		backwards: q.backwards.Push(t),
	}, nil
}

func (q NonFullQueue[T]) Dequeue() (Queue[T], error) {
	f, err := q.forwards.Pop()

	if !f.IsEmpty() {
		return NonFullQueue[T]{
			capacity:  q.capacity,
			forwards:  f,
			backwards: q.backwards,
		}, err
	}

	if !q.backwards.IsEmpty() {
		return NonFullQueue[T]{
			capacity:  q.capacity,
			forwards:  q.backwards.Reverse(),
			backwards: stack.EmptyStack[T]{},
		}, err
	}

	return EmptyQueue[T]{}, err
}

func (q NonFullQueue[T]) Peek() (T, error) {
	return q.forwards.Top()
}

func (q NonFullQueue[T]) IsEmpty() bool {
	return false
}

func (q NonFullQueue[T]) Length() uint {
	return q.forwards.Depth() + q.backwards.Depth()
}

func (q NonFullQueue[T]) Capacity() uint {
	return q.capacity
}

func (q NonFullQueue[T]) String() string {
	return fmt.Sprintf("fwd: %v\nbwd: %v", q.forwards, q.backwards)
}

type FullQueue[T any] struct {
	capacity            uint
	forwards, backwards stack.Stack[T]
}

func (q FullQueue[T]) Enqueue(t T) (Queue[T], error) {
	return q, errors.New("cannot enqueue at full capacity")
}

func (q FullQueue[T]) Dequeue() (Queue[T], error) {
	f, err := q.forwards.Pop()

	if !f.IsEmpty() {
		return NonFullQueue[T]{
			capacity:  q.capacity,
			forwards:  f,
			backwards: q.backwards,
		}, err
	}

	if !q.backwards.IsEmpty() {
		return NonFullQueue[T]{
			capacity:  q.capacity,
			forwards:  q.backwards.Reverse(),
			backwards: stack.EmptyStack[T]{},
		}, err
	}

	return EmptyQueue[T]{}, err
}

func (q FullQueue[T]) Peek() (T, error) {
	return q.forwards.Top()
}

func (q FullQueue[T]) IsEmpty() bool {
	return false
}

func (q FullQueue[T]) Length() uint {
	return q.forwards.Depth()
}

func (q FullQueue[T]) Capacity() uint {
	return q.capacity
}

func (q FullQueue[T]) String() string {
	return fmt.Sprintf("fwd: %v\nbwd: %v", q.forwards, q.backwards)
}
