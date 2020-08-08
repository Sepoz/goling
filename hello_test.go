// for testing a go file needs to be named like ..._test.go

package main

import "testing"

/* the function MUST start with the word Test
the function takes only one argument t *testing.T
t of the type *testing.T is like a "hook" to the testing library */
func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello Chris"

	if got != want {
		// %q wraps the value in a double quoted string
		t.Errorf("got %q want %q", got, want)
	}
}
