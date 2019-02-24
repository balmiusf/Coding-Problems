/*
Given a string of parentheses, write a function to compute the minimum number of
 parentheses to be removed to make the string valid (i.e. each open parenthesis is eventually closed).

For example, given the string "()())()", you should return 1. Given the string ")(",
you should return 2, since we must remove all of them.
*/

package main

import (
	"fmt"
)

func problem2(str string) int {
	outOfPlace := 0 // if there is a ')' before a '(' which can be identified by the total var.
	total := 0      // number of possible matches, 0 means everything is correct

	for _, char := range str {
		if char == '(' {
			total++ // possible pair
		} else {
			total--        // possibly found a match
			if total < 0 { // found a ')' out of place == ')' before '('
				outOfPlace++
				total = 0 // reset counter so it doesn't influence ')' before '(' for '(' before ')'
			}

		}
	}

	return outOfPlace + total
}

func main() {
	fmt.Println(problem2("()())()"))
	fmt.Println(problem2(")))((("))
	fmt.Println(problem2(")))((()))"))
	fmt.Println(problem2(")("))
}
