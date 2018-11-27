/*
There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.
Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.

For example, if N is 4, then there are 5 unique ways:

1, 1, 1, 1
2, 1, 1
1, 2, 1
1, 1, 2
2, 2

What if, instead of being able to climb 1 or 2 steps at a time, you could climb any number from a set of positive integers X?
For example, if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time.
*/

package main

import "fmt"

// assuming numberOfPossibleSteps[i] > 0
func waysToClimbStaircase(numberOfStairs int, possibleSteps []int) int {
	if numberOfStairs <= 0 {
		return 0
	}

	possibleWays := 0
	for _, step := range possibleSteps {
		if numberOfStairs-step == 0 {
			possibleWays++
		}
	}
	for _, step := range possibleSteps {
		possibleWays += waysToClimbStaircase(numberOfStairs-step, possibleSteps)
	}

	return possibleWays
}

func main() {
	possibleSteps := []int{1, 3, 5}
	n := 5
	fmt.Println(waysToClimbStaircase(n, possibleSteps))
}
