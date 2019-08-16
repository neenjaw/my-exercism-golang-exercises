/*
Package twofer is a library which contains a function that prints string response
*/
package twofer

import "fmt"

// ShareWith takes a string parameter, return a string to the form: "One for <name>, one for me."
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
