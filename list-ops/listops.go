// Package listops contains methods for operations on a list of integers
// represented by a slice of int
package listops

// IntList a type to represent a list of intergers
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(n int) int

// Foldl from the left reduce
func (l *IntList) Foldl(f binFunc, initial int) int {
	acc := initial

	for _, e := range *l {
		acc = f(acc, e)
	}

	return acc
}

// Foldr from the right reduce list
func (l *IntList) Foldr(f binFunc, initial int) int {
	acc := initial

	for i := range *l {
		acc = f((*l)[len(*l)-1-i], acc)
	}

	return acc
}

// Filter remove elements from list when the passed func evaluates to false
func (l *IntList) Filter(f predFunc) IntList {
	filtered := make(IntList, 0)

	for _, e := range *l {
		if f(e) {
			filtered = append(filtered, e)
		}
	}

	return filtered
}

// Length returns the length of the IntList
func (l *IntList) Length() int {
	count := 0

	for range *l {
		count++
	}

	return count
}

// Map applies a function to each element of the IntList
func (l *IntList) Map(f unaryFunc) IntList {
	mapped := make(IntList, len(*l))

	for i, e := range *l {
		mapped[i] = f(e)
	}

	return mapped
}

// Reverse returns a new IntList with the elements reversed
func (l *IntList) Reverse() IntList {
	reversed := make(IntList, len(*l))

	for i, e := range *l {
		reversed[len(*l)-1-i] = e
	}

	return reversed
}

// Append an IntList to another IntList
func (l *IntList) Append(a IntList) IntList {
	appended := make(IntList, len(*l)+len(a))

	for i, e := range *l {
		appended[i] = e
	}

	for i, e := range a {
		appended[i+len(*l)] = e
	}

	return appended
}

// Concat take a slice of IntList and return a new IntList with all elements appened
func (l *IntList) Concat(ll []IntList) IntList {
	concatted := make(IntList, len(*l))

	for i, e := range *l {
		concatted[i] = e
	}

	for _, iList := range ll {
		concatted = concatted.Append(iList)
	}

	return concatted
}
