package main

import "testing"

func TestHello(t *testing.T) {

	/**
	* 1) Write test
	* 2) Make compiler pass it
	* 3) Run test, see if it fails/passes. Is the error message meaningful?
	* 4) Write enough code to make test pass?
	* 5) Refactor.
	**/

	// To avoid duplicated code:
	assertCorrectMessage := func(t *testing.T, actual, expected string) {
		t.Helper() // needed to tell suite this is a helper!
		if expected != actual {
			t.Errorf("Expected %q, got %q", expected, actual)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		expected := "Hello, Pedro"
		actual := HelloWorld("Pedro")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, World when no string", func(t *testing.T) {
		expected := "Hello, World"
		actual := HelloWorld("")
		assertCorrectMessage(t, expected, actual)
	})
}
