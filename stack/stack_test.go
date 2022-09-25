package stack_test

import (
	"fmt"
	"testing"

	"framagit.org/attaboy/go-imm/stack"
)

func ExampleStack() {
	s := stack.NewStack[string]()
	s = s.Push("Hello")
	s = s.Push("World")
	s, _ = s.Pop()
	s, _ = s.Pop()
	_, err := s.Pop()
	fmt.Println(err)
	// Output: cannot pop empty stack
}

func TestStack(t *testing.T) {
	s := stack.NewStack[int]()
	if s.Depth() != 0 {
		t.Error("expected stack to have depth 0")
	}

	s1 := s.Push(52)
	if s1.Depth() == s.Depth() {
		t.Error("expected stack to be immutable")
	}

	s1 = s1.Push(666)
	top, err := s1.Top()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	if top != 666 {
		t.Errorf("wanted 666, got: %d", top)
	}

	fmt.Println(s1)

	ps, err := s1.Pop()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	pTop, err := ps.Top()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	if pTop != 52 {
		t.Errorf("wanted 52, got: %d", pTop)
	}

	ps, err = ps.Pop()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	ps, err = ps.Pop()
	if err == nil {
		t.Errorf("popping empty stack should result in err")
	}
}

func TestReferences(t *testing.T) {
	s1 := stack.NewStack[float32]()
	s2 := s1.Push(6.)
	s3 := s2.Push(.75)

	ps2, err := s3.Pop()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	if ps2 != s2 {
		t.Error("invalid references")
	}

	ps1, err := ps2.Pop()
	if err != nil {
		t.Errorf("wanted nil, got: %v", err)
	}
	if ps1 != s1 {
		t.Error("invalid references")
	}

	_, err = ps1.Pop()
	if err == nil {
		t.Error("popping empty stack should result in err")
	}
}

func TestReverse(t *testing.T) {
	s := stack.NewStack[rune]()
	fwd := "Hello, World!"
	for _, r := range fwd {
		s = s.Push(r)
	}

	s = s.Reverse()
	for _, r := range fwd {
		val, err := s.Top()
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
		if val != r {
			t.Errorf("wanted %c, got: %c", r, val)
		}

		s, err = s.Pop()
		if err != nil {
			t.Errorf("wanted nil, got: %v", err)
		}
	}
}
