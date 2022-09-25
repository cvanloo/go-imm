package queue_test

import (
	"fmt"
	"testing"

	"framagit.org/attaboy/go-imm/queue"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue[int](10)

	q, err := q.Enqueue(5)
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	q, err = q.Enqueue(6)
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}

	el, err := q.Peek()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	if el != 5 {
		t.Errorf("wanted 5, got: %d", el)
	}
	q, err = q.Dequeue()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}

	el, err = q.Peek()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	if el != 6 {
		t.Errorf("wanted 5, got: %d", el)
	}
	q, err = q.Dequeue()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty")
	}
}

func TestFullQueue(t *testing.T) {
	q := queue.NewQueue[int](5)

	vals := []int{1, 2, 3, 4, 5}

	var err error
	for _, v := range vals {
		q, err = q.Enqueue(v)
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
	}

	q, err = q.Enqueue(6)
	if err == nil {
		t.Error("expected exceeding capacity error, but got nil")
	}

	for _, v := range vals {
		val, err := q.Peek()
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
		if val != v {
			t.Errorf("wanted %d, got: %d", v, val)
		}
		q, err = q.Dequeue()
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
	}

	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}
}

func TestFullQueueWithDequeueInBetween(t *testing.T) {
	q := queue.NewQueue[int](5)

	vals := []int{1, 2, 3, 4}

	var err error
	for _, v := range vals {
		q, err = q.Enqueue(v)
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
	}
	fmt.Printf("%T\n", q)
	fmt.Println(q)
	// NonFullQueue
	// fwd: 1, nil
	// bwd: 4, 3, 2, nil

	q, _ = q.Dequeue()
	fmt.Printf("%T\n", q)
	fmt.Println(q)
	// NonFullQueue
	// fwd: 2, 3, 4, nil
	// bwd: nil

	q, _ = q.Enqueue(5)
	q, _ = q.Enqueue(6)
	fmt.Printf("%T\n", q)
	fmt.Println(q)
	// FullQueue
	// fwd: 2, 3, 4, nil
	// bwd: 6, 5, nil

	q, err = q.Enqueue(7)
	if err == nil {
		t.Error("expected exceeding capacity error, but got nil")
	}

	newVals := []int{2, 3, 4, 5, 6}
	for _, v := range newVals {
		val, err := q.Peek()
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
		if val != v {
			t.Errorf("wanted %d, got: %d", v, val)
		}
		q, err = q.Dequeue()
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
	}

	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}
}
