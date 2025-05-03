package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("say 'Hello' to people", func(t *testing.T) {
		got := Hello("Alexey", "")
		want := "Hello, Alexey!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when got empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello' in another language ", func(t *testing.T) {
		got := Hello("Alexey", "Spanish")
		want := "Hola, Alexey!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %s", got, want)
	}
}
