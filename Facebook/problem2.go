/*
Given an array of integers, write a function to determine whether the array could become non-decreasing by modifying at most 1 element.

For example, given the array [10, 5, 7], you should return true, since we can modify the 10 into a 1 to make the array non-decreasing.

Given the array [10, 5, 1], you should return false, since we can't modify any one element to get a non-decreasing array.

*/

package main

import "fmt"

func problem2(array []int) bool {
	if len(array) == 0 {
		return false
	}
	if len(array) == 1 || len(array) == 2 {
		return true
	}

	anyMoreTries := true
	for index := 0; index < len(array)-1; index++ {

		if anyMoreTries && array[index] > array[index+1] {
			anyMoreTries = false
		} else if array[index] > array[index+1] {
			return false
		}
	}

	return true
}

func main() {
	a := []int{10, 5, 6}
	fmt.Println(problem2(a))
}
