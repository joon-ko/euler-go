package solution

import "testing"

func TestSolution(t *testing.T) {

	assertEqual := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("numbers below 10", func(t *testing.T) {
		got := Solution(10)
		want := 23
		assertEqual(t, got, want)
	})

	t.Run("numbers below 1000", func(t *testing.T) {
		got := Solution(1000)
		want := 233168
		assertEqual(t, got, want)
	})
}
