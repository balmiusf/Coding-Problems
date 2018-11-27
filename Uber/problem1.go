/*
Given an array of integers,
return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5],
the expected output would be [120, 60, 40, 30, 24].
If our input was [3, 2, 1], the expected output would be [2, 3, 6].

*/
package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func uberExercise(arrayOfIntegers []int) {

	if len(arrayOfIntegers) == 0 {
		fmt.Println("[]")
		return
	} else if len(arrayOfIntegers) == 1 {
		fmt.Println("[ 0 ]")
		return
	}

	productTotalOfIndex := 1
	var productTotalList []int
	for i, _ := range arrayOfIntegers {
		for j, element := range arrayOfIntegers {
			if i == j {
				continue
			}
			productTotalOfIndex = productTotalOfIndex * element
		}
		productTotalList = append(productTotalList, productTotalOfIndex)
		productTotalOfIndex = 1
	}

	var buffer bytes.Buffer

	buffer.WriteString("\n[ ")
	buffer.WriteString(strconv.Itoa(productTotalList[0]))

	for index := 1; index < len(productTotalList); index++ {
		buffer.WriteString(", ")
		buffer.WriteString(strconv.Itoa(productTotalList[index]))
	}

	buffer.WriteString(" ]\n")

	fmt.Println(buffer.String())
}

func main() {
	a := []int{1, 2, 3, 4, 5}

	typeOfArg := reflect.TypeOf(a)
	if reflect.Slice == typeOfArg.Kind() {
		uberExercise(a)
	} else {
		return
	}

}
