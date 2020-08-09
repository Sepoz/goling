// for testing a go file needs to be named like ..._test.go

package main

import "testing"

/* the function MUST start with the word Test
the function takes only one argument t *testing.T
t of the type *testing.T is like a "hook" to the testing library */
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty strin is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}
