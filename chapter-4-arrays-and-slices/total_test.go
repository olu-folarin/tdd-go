package total

import "testing"

func TestTotal(t *testing.T) {
	t.Run("collection of 6 integers", func(t *testing.T) {
		sum := []int{1,3,5, 7, 9, 11}
		got := Total(sum)
		want := 36

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		sum := []int{1,4,6,8,9,2,45,5,6,7,8}
		got := Total([]int(sum))
		want := 101

		if got != want {
			t.Errorf("got %d want %d give %v", got, want, sum)
		}
	})
}