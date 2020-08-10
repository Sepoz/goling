package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

/* when the benchmark code is executed, it runs b.N times and measures how long it takes
To run the benchmarks do (go test -bench=.)
136 ns/op means is our function takes on average 136 nanoseconds to run */
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
