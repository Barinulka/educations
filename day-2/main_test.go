package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAdd(t *testing.T) {
	got := Add(3, 5)
	want := 8

	assertCorrectMessage(t, got, want)
}
