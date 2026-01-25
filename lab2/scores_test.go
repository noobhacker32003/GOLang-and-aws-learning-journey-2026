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


// func TestCalculateTotalSlice(t *testing.T) {
// 	t.Run("with 5 scores", func(t *testing.T) {
// 		scores := []int{10, 15, 20, 12, 18}
// 		got := CalculateTotal(scores)
// 		want := 75
		
// 		if got != want {
// 			t.Errorf("got %d but wanted %d", got, want)
// 		}
// 	})
	
// 	t.Run("with 3 scores", func(t *testing.T) {
// 		scores := []int{25, 30, 35}
// 		got := CalculateTotal(scores)
// 		want := 90
		
// 		if got != want {
// 			t.Errorf("got %d but wanted %d", got, want)
// 		}
// 	})
// }



// t.Run("with empty slice", func(t *testing.T) {
// 	scores := []int{}
// 	got := CalculateTotal(scores)
// 	want := 0
	
// 	if got != want {
// 		t.Errorf("got %d but wanted %d", got, want)
// 	}
// })

// t.Run("with single score", func(t *testing.T) {
// 	scores := []int{42}
// 	got := CalculateTotal(scores)
// 	want := 42
	
// 	if got != want {
// 		t.Errorf("got %d but wanted %d", got, want)
// 	}
// })
// }

func TestCalculateAverages(t *testing.T) {
	player1 := []int{}
	player2 := []int{25, 35}
	player3 := []int{5, 15, 25, 35}		

	got := CalculateAverages(player1, player2, player3)
	want := []int{0, 30, 20}

	if !equal(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}

}
func TestCalculateTotals(t *testing.T) {
	player1 := []int{10, 15, 20}
	player2 := []int{25, 30}
	player3 := []int{5, 10, 15, 20}
	
	got := CalculateTotals(player1, player2, player3)
	want := []int{45, 55, 50}
	
	if !equal(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}

	t.Run("single player", func(t *testing.T) {
	player1 := []int{100, 200, 300}
	
	got := CalculateTotals(player1)
	want := []int{600}
	
	if !equal(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}



})

t.Run("no players", func(t *testing.T) {
	got := CalculateTotals()
	want := []int{}
	
	if len(got) != len(want) {
		t.Errorf("got %v but wanted empty slice", got)
	}
})
	
}

// Helper function to compare slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBestRounds(t *testing.T) {
	t.Run("normal cases", func(t *testing.T) {
		player1 := []int{5, 10, 15, 20}   // Warm-up: 5, best: 20
		player2 := []int{30, 25, 35}      // Warm-up: 30, best: 35
		
		got := BestRounds(player1, player2)
		want := []int{20, 35}
		
		if !equal(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})


	t.Run("edge case: player with only warm-up", func(t *testing.T) {
	player1 := []int{10}  // Only warm-up, no actual rounds
	
	got := BestRounds(player1)
	want := []int{0}
	
	if !equal(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
	})

	t.Run("edge case: empty slice", func(t *testing.T) {
		player1 := []int{}
		
		got := BestRounds(player1)
		want := []int{0}
		
		if !equal(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})

}
