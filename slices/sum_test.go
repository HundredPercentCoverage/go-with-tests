package slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	want := []int{6, 22}
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6, 7})

	if !slices.Equal(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSumTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !slices.Equal(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
	t.Run("sum of tails of two slices", func(t *testing.T) {
		got := SumTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})

	t.Run("safely handle empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}
