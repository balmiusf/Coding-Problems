/*
uppose you have a multiplication table that is N by N. That is, a 2D array where the value at the i-th row and j-th column is (i + 1) * (j + 1) (if 0-indexed) or i * j (if 1-indexed).

Given integers N and X, write a function that returns the number of times X appears as a value in an N by N multiplication table.

For example, given N = 6 and X = 12, you should return 4, since the multiplication table looks like this:

| 1 | 2 | 3 | 4 | 5 | 6 |

| 2 | 4 | 6 | 8 | 10 | 12 |

| 3 | 6 | 9 | 12 | 15 | 18 |

| 4 | 8 | 12 | 16 | 20 | 24 |

| 5 | 10 | 15 | 20 | 25 | 30 |

| 6 | 12 | 18 | 24 | 30 | 36 |

And there are 4 12's in the table.


*/

package main

import (
	"fmt"
	"math/big"
)

func isPrime(possiblePrime int) bool {
	prime := new(big.Int)

	return prime.ProbablyPrime(possiblePrime)
}

func problem2(n int, x int) (xtimes int) {
	if x < 0 || n < 0 || (n == 0 && x > 0) || (n*n < x) {
		return 0
	}

	if (n == 0 && x == 0) || n*n == x {
		return 1
	}

	if isPrime(x) && x <= n {
		return 2
	}

	for factor := 1; factor <= n; factor++ {

		// while the mult factor is to low -> ignore next steps and increase factor
		if factor*n < x {
			continue
		}

		for v := 1; v <= n; v++ {
			if v*factor == x {
				xtimes++
			}
		}
	}

	return xtimes
}

func main() {
	fmt.Println(problem2(32, 31))
}
