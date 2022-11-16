package week1

// BreakingRecords - Maria plays college basketball and wants to go pro.
// Given the scores for a season, determine the number of times Maria breaks her records for most and least points scored during the season.
// 11/16/22 3:41 - 3:56; 15 minutes, like two for the basic solution, three or four minutes getting distracted, and 8 minutes generating tests for the "don't count the first game" fix.
// TAKEAWAYS:
//  * Watch out for Go's default values. This was multiple takes on an off-by-one error.
func BreakingRecords(scores []int32) []int32 {
	var minCount, maxCount int32
	currMax := scores[0]
	currMin := scores[0]
	for _, s := range scores {
		if s > currMax {
			currMax = s
			maxCount++
		}
		if s < currMin {
			currMin = s
			minCount++
		}
	}
	return []int32{maxCount, minCount} // off-by-one for first iteration
}
