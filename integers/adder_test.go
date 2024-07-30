package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(4, 2)
	want := 6

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(4, 4)
	fmt.Println(sum)
	// Output: 8
}
