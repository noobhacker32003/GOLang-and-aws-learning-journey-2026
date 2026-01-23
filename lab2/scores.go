// package tournament

// func CalculateTotal(scores [5]int) int {
// 	total := 0
// 	for i := 0; i < 5; i++ {
// 		total = total + scores[i]
// 	}
// 	return total
// }

package tournament

func CalculateTotal(scores [5]int) int {
	total := 0
	for _, score := range scores {
		total += score
	}
	return total
}
