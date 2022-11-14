package week1

import (
	"sort"
)

// MinMaxSum - given five positive integers, find the minimum and maximum values that can be calculated by summing
// exactly four of the five integers.
// Started at 11:18AM, finished at 11:29AM including tests. 11 minutes.
// HackerRank seems to have a thing for int32. They'll be the same on most systems.
// HackerRank also doesn't let you use `x` libraries. Changing to standard library.
// Looked at the editorial - they have a cute solution for doing this in a single pass.
// Basically, you find the total sum, min value, and max value over one pass, then return sum - max and sum - min.
// TAKEAWAYS:
// * Okay, I missed overflow conditions. This wouldn't have happened if we just got to use int, since that would've been an int64 on my system.
// * I also mixed up the sort direction on sort.Slice - if you preserve the index order, your sort will follow the operand, e.g. "<" -> smallest to largest
// * Cute solutions to avoid multiple passes seem to be encouraged in these problems. Which is irritating, because who knows if they're even more efficient, and they usually aren't more clear and maintainable.
// * Including playing with sorting and types, fully correct on HR at 11:38, so 20 minutes.
func MinMaxSum(arr []int32) (min int64, max int64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := range arr {
		// Sum first four of sorted.
		if i < len(arr)-1 {
			min += int64(arr[i])
		}
		// Sum last four of sorted.
		if i > 0 {
			max += int64(arr[i])
		}
	}
	return min, max
}

// MinMaxSum2 is the single-pass solution contained in the HackerRank editorial. I want to see how much faster it is.
// Result - it's much faster. It's not hard to read either.
func MinMaxSum2(arr []int32) (int64, int64) {
	minV, maxV := arr[0], arr[0]
	var sum int64
	for i := range arr {
		sum += int64(arr[i])
		if arr[i] > maxV {
			maxV = arr[i]
		}
		if arr[i] < minV {
			minV = arr[i]
		}
	}
	return sum - int64(maxV), sum - int64(minV)
}
