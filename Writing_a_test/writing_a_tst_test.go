package main

import(
	"testing"
)

// The first letter after the word "Test" has to be capitalized.
/* Type T object is the object that is given by Go's test runner,
that the program that's actually executing our tests. That test runner is going to pass us this T object.
We use that T object to communicate back to the test runner, about what's going on inside our tests.
That's why we have a pointer, we're sharing memory back with the test runner.*/
func TestAdd(t *testing.T) {
	// arrange, set up test conditions
	l, r := 1, 2
	expect := 3
	// action
	got := Add(l, r)

	// assert
	if expect != got {
		t.Errorf("Failed to add %v and %v. Got %v, expected %v\n", l, r, got, expect)
	}
}