package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello world' if empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Cherie", "French")
		want := "Bonjour Cherie"
		assertCorrectMessage(t, got, want)
	})
}

// Helper function
// testing.TB is an interface that both *testing.T and *testing.B implement
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() tells the test suite that this is a helper - affects the line number shown on fails
	// i.e. if it fails the line shown will be the call to assertCorrectMessage rather than the t.Errorf() call below
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
