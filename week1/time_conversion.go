package week1

import (
	"fmt"
	"strconv"
)

// TimeConversion - Given a time in -hour AM/PM format, convert it to military (24-hour) time.
// Started at 11:23AM, 11/15/2022. Finished at 11:48AM. 25 minutes.
// Sample input: "07:05:45PM"
// We can grab the second to last character in constant time. If it's an A, leave the time; P -> add 12 hours to the first two characters as integers.
// Special cases - basically just the 12s. There might be more, I'll add cases.
//	 - 12:20PM -> it's a P, but don't add anything
// 	 - 12:20AM -> it's an A, but subtract 12.
//
// When converting zeroes, need to preserve the leading zero.
// This is basically constant time because of direct slice access and taking advantage of the input structure. Constraints help!
// TAKEAWAYS:
// * Strings vs. runes, slicing vs. indexing -- gotta be precise.
// * Are there better ways to do these conversions? Maybe I don't even need to treat it as an integer, since the runes are already integers. That would eliminate the strconv calls.
// * fmt.Sprintf is best for formatting numbers. I reached for a strings.Builder for a sec, no no!
func TimeConversion(s string) string {
	isAM := s[len(s)-2] == 'A'
	is12 := s[:2] == "12"
	hour, _ := strconv.Atoi(s[:2]) // ignoring Atoi error
	switch {
	case isAM && is12:
		hour -= 12
	case !isAM && !is12:
		hour += 12
	}
	return fmt.Sprintf("%02s%s", strconv.Itoa(hour), s[2:len(s)-2])
}
