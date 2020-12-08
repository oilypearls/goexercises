package main

import "testing"

func TestAdding(t *testing.T) {
	sum := Adding(2, 6)
	expected := 8

	if sum != expected {
		t.Errorf("Expected %d, but got %d", sum, expected)
	}
}
