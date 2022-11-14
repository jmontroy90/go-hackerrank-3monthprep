package week1

import (
	"fmt"
)

// PlusMinus - took about 20 minutes, most of it looking up different strconv functionality which I didn't actually need.
// Since they want formatted strings with 6 decimals of precision, we'll return strings here.
// TAKEAWAYS:
// * Spend some more time with the strconv library.
// * Type conversions in general are annoying on these HackerRank questions.
func PlusMinus(arr []int32) (pos string, neg string, zero string) {
	var nPos, nNeg, nZero int
	for _, v := range arr {
		switch {
		case v < 0:
			nNeg++
		case v > 0:
			nPos++
		default:
			nZero++
		}
	}
	tlen := len(arr)
	return formattedDivide(nPos, tlen), formattedDivide(nNeg, tlen), formattedDivide(nZero, tlen)
}

func formattedDivide(i1, i2 int) string {
	return fmt.Sprintf("%1.6f", float64(i1)/float64(i2))
}
