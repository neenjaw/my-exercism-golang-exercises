// Package collatzconjecture contains functions pertaining to the collatz conjecture (3n + 1 problem)
package collatzconjecture

import "errors"

// CollatzConjecture takes and integer and returns the steps
//   taken to reach 1 by the following algorithm:
//   if even: n = n / 2
//   if odd: n = 3n + 1
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("number must be greater than zero")
	}

	steps := 0

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		steps++
	}

	return steps, nil
}
