/*
This problem was asked by Apple.

Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.

*/

package main

import (
	"fmt"
	"reflect"
	"time"
)

func doNothing() {
	return
}

func sum(a int, b string) int {
	fmt.Println(b)
	return a
}

func jobScheduler(f interface{}, milliseconds int64, args ...interface{}) error {

	if milliseconds < 0 {
		return fmt.Errorf("%d < 0", milliseconds)
	}
	// check if function is a function
	function := reflect.ValueOf(f)
	if function.Kind() != reflect.Func {
		return fmt.Errorf("Function '%s' is not defined", f)
	}

	if len(args) != function.Type().NumIn() {
		return fmt.Errorf("Args provided: %d; Args needed: %d", len(args), function.Type().NumIn())
	}

	newArgs := make([]reflect.Value, len(args))

	functionType := function.Type()

	for i, arg := range args {

		// check if provided arg type matches the actual function arg type
		if functionType.In(i).Kind() == reflect.ValueOf(arg).Kind() {
			newArgs[i] = reflect.ValueOf(arg)
		} else {
			return fmt.Errorf("provided arg at position '%d' does not match the type of actual function arg", i)
		}

	}

	time.Sleep(time.Duration(milliseconds)) // wait to be executed

	result := function.Call(newArgs)
	if len(result) >= 1 {
		// test if called correctly
		for _, value := range result {
			fmt.Println(value)
		}
	}

	return nil
}

func main() {
	fmt.Println(jobScheduler(sum, 120, 1, 2))
}
