package fizbuz

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFizBuz(t *testing.T) {
	t.Run("Test It Return FizBuzz", func(t *testing.T) {
		got := FizBuz(15)
		want := "FizzBuzz"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Test It Return Fizz When i%3", func(t *testing.T) {
		got := FizBuz(3)
		want := "Fizz"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Test It Return Buzz When i%5", func(t *testing.T) {
		got := FizBuz(5)
		want := "Buzz"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Test It Return number", func(t *testing.T) {
		got := FizBuz(2)
		want := "2"

		assertCorrectMessage(t, got, want)
	})
}
