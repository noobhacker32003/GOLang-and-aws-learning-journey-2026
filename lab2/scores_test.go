package tournament

import "testing"

// func TestCalculateTotal(t *testing.T) {
// 	scores := [5]int{10, 15, 20, 12, 18}
// 	got := CalculateTotal(scores)
// 	want := 75

// 	if got != want {
// 		t.Errorf("got %d but wanted %d", got, want)
// 	}
// }


func TestCalculateTotalSlice(t *testing.T) {
	t.Run("with 5 scores", func(t *testing.T) {
		scores := []int{10, 15, 20, 12, 18}
		got := CalculateTotal(scores)
		want := 75
		
		if got != want {
			t.Errorf("got %d but wanted %d", got, want)
		}
	})
	
	t.Run("with 3 scores", func(t *testing.T) {
		scores := []int{25, 30, 35}
		got := CalculateTotal(scores)
		want := 90
		
		if got != want {
			t.Errorf("got %d but wanted %d", got, want)
		}
	})
}
