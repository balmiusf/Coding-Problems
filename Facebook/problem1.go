/*
Given a list of integers, return the largest product that can be made by multiplying any three integers.

For example, if the list is [-10, -10, 5, 2], we should return 500, since that's -10 * -10 * 5.

You can assume the list has at least three integers.
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Simple way of solving would be to sort the array first
then compare the product of the first two elements (possibly highest negative numbers)
and the highest number in the list (lowest negative or highest positive)
with the product of the last three elements in the array
*/

func problem1(set []int) (product int) {
	sort.Ints(set)
	size := len(set)
	return int(math.Max(float64(set[0]*set[1]*set[size-1]), float64(set[size-1]*set[size-2]*set[size-3])))
}

func main() {
	a := []int{-10, -10, 5, 2}
	fmt.Println(problem1(a))
}
