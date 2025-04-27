package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Alexey")
	want := "Hello, Alexey"

	if (got != want) {
		t.Errorf("got %q want %s", got, want)
	}
}