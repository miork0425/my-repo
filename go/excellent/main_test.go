package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Odd" {
		t.Errorf("expected: odd, actual: %s", result)
	}
}