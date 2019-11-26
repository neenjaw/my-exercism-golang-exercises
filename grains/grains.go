/*
Package grains is a package of functions that calculate the number of grains of
wheat on a chessboard given that the number on each square doubles.

As the tale goes, "There once was a wise servant who saved the life of a prince.
The king promised to pay whatever the servant could dream up. Knowing that the
king loved chess, the servant told the king he would like to have grains of wheat.
One grain on the first square of a chess board, with the number of grains doubling
on each successive square.""

There are 64 squares on a chessboard (where square 1 has one grain, square 2 has
two grains, and so on).
*/
package grains

import "errors"

// Square takes an number (n) and shifts an initial value of 1 by that number
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid number")
	}

	return 1 << (n - 1), nil
}

// Total returns the sum, starting with 1, then the result of 1 shifted left on
// each iteration
func Total() uint64 {
	return 1<<64 - 1
}
