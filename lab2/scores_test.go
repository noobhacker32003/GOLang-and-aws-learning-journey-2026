package tournament

import "testing"

func TestCalculateTotal(t *testing.T) {
	scores := [5]int{10, 15, 20, 12, 18}
	got := CalculateTotal(scores)
	want := 75

	if got != want {
		t.Errorf("got %d but wanted %d", got, want)
	}
}
