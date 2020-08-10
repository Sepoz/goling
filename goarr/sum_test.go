package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func TestMultiply(t *testing.T) {

	numbers := []int{2, 5}

	got := Multiply(numbers)
	want := 10

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {

	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{2, 5, 6, 8, 9}

	for i := 0; i < b.N; i++ {
		SumAllTails(numbers1, numbers2)
	}
}
