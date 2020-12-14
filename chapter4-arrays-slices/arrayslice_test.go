package main

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, actual int) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	}

	t.Run("Sum up some easy ass numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15
		actual := Sum(numbers)
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Add up all 0s", func(t *testing.T) {
		numbers := []int{0, 0, 0, 0, 0}
		expected := 0
		actual := Sum(numbers)
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Test an array with 4 length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		expected := 10
		actual := Sum(numbers)
		assertCorrectMessage(t, expected, actual)
	})
}
