package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPrintInfo(t *testing.T) {
	got := PrintInfo("Алексей", 32, "Изучает GO")
	want := "Имя: Алексей, возраст: 32, статус: 'Изучает GO'"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
