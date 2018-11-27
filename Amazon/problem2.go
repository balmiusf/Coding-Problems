/*
Given an integer k and a string s,
find the length of the longest substring that contains at most k distinct characters.

For example, given s = "abcba" and k = 2,
the longest substring with k distinct characters is "bcb".
*/

package main

import (
	"fmt"
	"strings"
)

func problemo(s string, k int) string {

	if len(s) == 0 || k <= 0 || len(s) < k {

		return ""
	}

	var temp strings.Builder
	numberOfDistinct := 1
	longestSubstring := string(s[0])

	// brute force approach
	for i, pos1 := range s {

		temp.WriteString(string(pos1))
		for j := i + 1; j < len(s); j++ {

			// if the char is not present in the temporary substring
			// it means a new distinct char was found
			if !strings.Contains(temp.String(), string(s[j])) {
				// does the found distinct char violates condition 'k' (maximum distinct characters)?
				if numberOfDistinct+1 > k {

					break
				}

				numberOfDistinct++
			}

			temp.WriteString(string(s[j]))

			if len(longestSubstring) < temp.Len() {

				longestSubstring = temp.String()
			}
		}

		numberOfDistinct = 1 // reset counter
		temp.Reset()
	}

	return longestSubstring
}

func main() {
	fmt.Println(problemo("abcba", 2))
}
