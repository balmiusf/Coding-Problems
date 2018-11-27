/*
This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear time and constant space.
In other words, find the lowest positive integer that does not exist in the array.
The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.
*/

package main

import (
	"fmt"
	"sort"
)

func missingNumbers(arrayOfNumbers []int) int {

	if len(arrayOfNumbers) == 0 {
		return 0
	}

	sort.Ints(arrayOfNumbers)
	temporary := 1 // used to find missing values

	for index := 0; index < len(arrayOfNumbers)-1; {

		if arrayOfNumbers[index]+temporary != arrayOfNumbers[index+1] && arrayOfNumbers[index] != arrayOfNumbers[index+1] {

			temporary++
			if arrayOfNumbers[index]+temporary-1 == 0 {
				continue
			}

			return arrayOfNumbers[index] + temporary - 1

		} else {
			index++
			temporary = 1 // reset
		}
	}

	lastPosition := len(arrayOfNumbers) - 1
	if arrayOfNumbers[lastPosition]+1 == 0 || arrayOfNumbers[lastPosition] < 0 {
		return 1
	} else {
		return arrayOfNumbers[lastPosition] + 1
	}
}

func main() {
	a := []int{3, 4, -1, 1, 1, 2}
	fmt.Println(missingNumbers(a))
}
