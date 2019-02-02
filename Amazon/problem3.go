/*
An sorted array of integers was rotated an unknown number of times.

Given such an array, find the index of the element in the array in faster than linear time.
If the element doesn't exist in the array, return null.

For example, given the array [13, 18, 25, 2, 8, 10] and the element 8, return 4 (the index of 8 in the array).

You can assume all the integers in the array are unique.

*/

package main

import "fmt"

// iterative implementation of binary search algorithm.
func binarySearch(array []int, left int, right int, v int) int {

	for left <= right {
		middle := left + (right-left)/2
		if array[middle] == v {
			return middle
		}

		if array[middle] < v {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func main() {
	a := []int{13, 18, 25, 2, 8, 10}

	index := binarySearch(a, 0, len(a)-1, 13)
	if index != -1 {
		fmt.Println(index)
	} else {
		fmt.Println(nil)
	}
}
