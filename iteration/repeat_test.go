package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// What is b.N?
	// Benchmark functions are ran several times by the testing package
	// b.N will increase each time until the benchmark runner is satisfied
	// with the stability of the benchmark
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 6)
	fmt.Println(result)
	// Output: aaaaaa
}
