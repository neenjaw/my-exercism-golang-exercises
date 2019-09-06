/*
Package strain contains functions to filter elements in a slice based on passed in function
*/
package strain

// Ints collection
type Ints []int

// Keep is a method for the type Ints, which accepts a function which, based on the boolean result,
// removes elements from the slice
func (i Ints) Keep(f func(int) bool) Ints {
	var kept Ints

	for _, e := range i {
		if f(e) {
			kept = append(kept, e)
		}
	}

	return kept
}

// Discard is a method for the type Ints, which accepts a function which, based on the boolean result,
// removes elements from the slice.  Since this is the inverse of the Keep method, the Keep method is
// called and the function is wrapped with an anonymous function returning the inverse of the result
func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(i int) bool { return !f(i) })
}

// Lists collection
type Lists [][]int

// Keep is a method for the type Lists, which accepts a function which, based on the boolean result,
// removes elements from the slice
func (l Lists) Keep(f func([]int) bool) Lists {
	var kept Lists

	for _, e := range l {
		if f(e) {
			kept = append(kept, e)
		}
	}

	return kept
}

// Strings collection
type Strings []string

// Keep is a method for the type Strings, which accepts a function which, based on the boolean result,
// removes elements from the slice
func (s Strings) Keep(f func(string) bool) Strings {
	var kept Strings

	for _, e := range s {
		if f(e) {
			kept = append(kept, e)
		}
	}

	return kept
}
