package week1

import (
	"sort"
)

// MinMaxSum - given five positive integers, find the minimum and maximum values that can be calculated by summing
// exactly four of the five integers.
// Started at 11:18AM, finished at 11:29AM including tests. 11 minutes.
// HackerRank seems to have a thing for int32. They'll be the same on most systems.
// HackerRank also doesn't let you use `x` libraries. Changing to standard library.
// RESULTS:
// * Okay, I missed overflow conditions. This wouldn't have happened if we just got to use int, since that would've been an int64 on my system.
// * I also mixed up the sort direction on sort.Slice - if you preserve the index order, your sort will follow the operand, e.g. "<" -> smallest to largest
// * Including playing with sorting and types, fully correct on HR at 11:38, so 20 minutes.
func MinMaxSum(arr []int32) (min int64, max int64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := range arr {
		if i < len(arr)-1 {
			min += int64(arr[i])
		}
		if i > 0 {
			max += int64(arr[i])
		}
	}
	return min, max
}
