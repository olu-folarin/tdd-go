package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// a test that returns the string "Hello, folarin!" if the string "folarin!" is passed
	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("folarin!", "")
		want := "Hello, folarin!"

		assertCorrectMessage(t, got, want)
	
		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
	})
	t.Run("say 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
	})
	t.Run("say Hello in Yoruba", func(t *testing.T) {
		got := Hello("ikechukwu!", "Yoruba")
		want := "Bawo ni, ikechukwu!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in Spanish", func(t *testing.T) {
		got := Hello("Lade!", "Spanish")
		want := "Hola, Lade!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in German", func(t *testing.T) {
		got := Hello("Deola!", "German")
		want := "Hallo, Deola!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// The test function is a function that takes a pointer to testing.T as an argument.
// The testing.T type is a type that is provided by the testing package.
// It is used to report whether a test has passed or failed.
// The t.Errorf function is used to report an error and continue the test.
// The t.Run function is used to group tests together.

// TESTING: Helper functions
// This reduces duplication and improves the readability of our tests. We need to pass in t *testing.T so that we can tell the test code to fail when we need to.
// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, or a benchmark?