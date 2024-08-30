package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	// This works because bytes.Buffer implements Write()
	// meaning that it adheres to the io.Writer interface
	// as does os.Stdout
	Greet(&buffer, "potato")

	got := buffer.String()
	want := "hello potato"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
