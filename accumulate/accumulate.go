/*
Package accumulate
*/
package accumulate

// Accumulate applied a converting function to each string supplied in the string slice
func Accumulate(strs []string, converter func(string) string) []string {
	var converted []string

	for _, str := range strs {
		converted = append(converted, converter(str))
	}

	return converted
}
