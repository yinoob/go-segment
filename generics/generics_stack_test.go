package main

import "testing"

func TestStackPopReturnsLastPushedValue(t *testing.T) {
	var stack Stack[int]

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	for _, want := range []int{30, 20, 10} {
		got, ok := stack.Pop()
		if !ok {
			t.Fatalf("Pop() ok = false, want true for value %d", want)
		}
		if got != want {
			t.Fatalf("Pop() = %d, want %d", got, want)
		}
	}
}

func TestStackPopEmptyReportsFalseAndZeroValue(t *testing.T) {
	var stack Stack[string]

	got, ok := stack.Pop()
	if ok {
		t.Fatal("Pop() ok = true, want false for empty stack")
	}
	if got != "" {
		t.Fatalf("Pop() = %q, want zero value", got)
	}
}
