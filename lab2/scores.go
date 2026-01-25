// package tournament

// func CalculateTotal(scores [5]int) int {
// 	total := 0
// 	for i := 0; i < 5; i++ {
// 		total = total + scores[i]
// 	}
// 	return total
// }

// package tournament

// func CalculateTotal(scores [5]int) int {
// 	total := 0
// 	for _, score := range scores {
// 		total += score
// 	}
// 	return total
// }

package tournament

func CalculateTotal(scores []int) int {
	total := 0
	for _, score := range scores {
		total += score
	}
	return total
}

func CalculateTotals(playerScores ...[]int) []int {
	totals := make([]int, len(playerScores))

	for i, scores := range playerScores {
		totals[i] = CalculateTotal(scores)
	}

	return totals
}

func CalculateAverages(playerScores ...[]int) []int {
	averages := make([]int, len(playerScores))
	for i, scores := range playerScores {
		if len(scores) == 0 {
			averages[i] = 0
		} else {
			total := CalculateTotal(scores)
			averages[i] = total / len(scores)
		}
	}

	return averages
}