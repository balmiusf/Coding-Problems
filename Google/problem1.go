/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

*/

package main

import "fmt"

func sumOfValues(arrayOfIntegers []int, k int) bool {

	if len(arrayOfIntegers) == 0 || len(arrayOfIntegers) == 1 {
		return false
	}

	var needed map[int]int
	needed = make(map[int]int)

	for _, element := range arrayOfIntegers {
		_, exists := needed[element]
		if exists {
			return true
		} else {
			needed[k-element] = element
			// fmt.Println(needed[k-element])
		}
	}

	return false
}

func main() {
	a := []int{17, 0, 15, 3, 7}
	fmt.Println(sumOfValues(a, 16))
}
