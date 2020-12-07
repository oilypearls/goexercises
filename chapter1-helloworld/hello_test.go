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

	t.Run("Say hello to people in English", func(t *testing.T) {
		expected := "Hello, Pedro"
		actual := HelloWorld("Pedro", "English")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, World when no string in English", func(t *testing.T) {
		expected := "Hello, World"
		actual := HelloWorld("", "English")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, World but in Spanish with a name", func(t *testing.T) {
		expected := "Hola, mi amor"
		actual := HelloWorld("mi amor", "Spanish")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, but in Spanish, no name", func(t *testing.T) {
		expected := "Hola, Mundo"
		actual := HelloWorld("", "Spanish")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, World but in French with a name", func(t *testing.T) {
		expected := "Bonjour, Antoine"
		actual := HelloWorld("Antoine", "French")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, but in French, no name", func(t *testing.T) {
		expected := "Bonjour, Monde"
		actual := HelloWorld("", "French")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, World but in German with a name", func(t *testing.T) {
		expected := "Hallo, Jochen"
		actual := HelloWorld("Jochen", "German")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello, but in French, no name", func(t *testing.T) {
		expected := "Hallo, Welt"
		actual := HelloWorld("", "German")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say Hello -> default to English no name none, func", func(t *testing.T) {
		expected := "Hello, World"
		actual := HelloWorld("", "")
		assertCorrectMessage(t, expected, actual)

	})

}
