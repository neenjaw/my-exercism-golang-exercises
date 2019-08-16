/*
Package proverb is a library which exports a function to build the classic proverb with an array of items:

Example input:

["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"]

Example output:

[
"For want of a nail the shoe was lost.",
"For want of a shoe the horse was lost.",
"For want of a horse the rider was lost.",
"For want of a rider the message was lost.",
"For want of a message the battle was lost.",
"For want of a battle the kingdom was lost.",
"And all for the want of a nail.",
]
*/
package proverb

import "fmt"

// Proverb accepts an array of strings representing items in the proverb,
// returning a slice with the lines of the proverb in at each index in order
func Proverb(rhymeItems []string) []string {
	if len(rhymeItems) == 0 {
		return []string{}
	}

	rhyme := make([]string, len(rhymeItems))
	firstItem := rhymeItems[0]

	for index := 0; index < len(rhymeItems)-1; index++ {
		rhyme[index] = fmt.Sprintf("For want of a %s the %s was lost.", rhymeItems[index], rhymeItems[index+1])
	}

	rhyme[len(rhymeItems)-1] = fmt.Sprintf("And all for the want of a %s.", firstItem)

	return rhyme
}
