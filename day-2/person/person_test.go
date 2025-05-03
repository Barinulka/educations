package main

import "testing"

func TestPerson(t *testing.T) {
	if !IsAdult(18) {
		t.Errorf("Все хорошо")
	}

	if IsAdult(17) {
		t.Errorf("Все плохо")
	}
}
