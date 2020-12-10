package main

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected %q, got %q", expected, actual)
		}
	}

	t.Run("When number is 0", func(t *testing.T) {
		expected := ""
		actual := Repeat("c", 0)
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Repeated itself 5  times", func(t *testing.T) {
		expected := "11111"
		actual := Repeat("1", 5)
		assertCorrectMessage(t, expected, actual)
	})
}

// We can use BENCHMARK(ING)!
// Performance testing, to keep it simple.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
