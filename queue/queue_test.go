package queue_test

import (
	"testing"

	"framagit.org/attaboy/go-imm/queue"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue[int]()

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
