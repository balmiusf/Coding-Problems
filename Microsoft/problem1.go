/*
A number is considered perfect if its digits sum up to exactly 10.

Given a positive integer n, return the n-th perfect number.

For example, given 1, you should return 19. Given 2, you should return 28.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// always assuming n to be >= 0
func problem1(n int) (perfectNumber int) {
	if n == 0 {
		return 10
	}
	if n == 10 {
		return 0
	}

	var str strings.Builder

	// single digit numbers
	str.WriteString(strconv.Itoa(n))

	if n < 10 {
		str.WriteString(strconv.Itoa(10 - n))
	} else {

		temp := n
		sum := 0
		for temp != 0 {
			sum += temp % 10
			temp /= 10
		}

		str.WriteString(strconv.Itoa(10 - sum))
	}
	perfectNumber, _ = strconv.Atoi(str.String())
	return perfectNumber
}

func main() {
	fmt.Println(problem1(19))
}
