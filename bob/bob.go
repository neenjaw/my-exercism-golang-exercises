// Package bob contains functions to allow bob to reply
package bob

import (
	"strings"
)

const (
	statementResponse string = "Whatever."
	shoutingResponse  string = "Whoa, chill out!"
	questionResponse  string = "Sure."
	forcefulResponse  string = "Calm down, I know what I'm doing!"
	silentResponse    string = "Fine. Be that way!"
)

// Hey takes a string as input, deciphers it, then returns his response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if isSilent(remark) {
		return silentResponse
	} else {

		letters := hasLetters(remark)
		question := isQuestion(remark)
		shout := isShout(remark)

		if letters && shout && question {
			return forcefulResponse
		} else if letters && shout {
			return shoutingResponse
		} else if question {
			return questionResponse
		}
	}

	return statementResponse
}

func hasLetters(remark string) bool {
	return strings.ToUpper(remark) != strings.ToLower(remark)
}

func isShout(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func isSilent(remark string) bool {
	return remark == ""
}

func isQuestion(remark string) bool {
	if len(remark) == 0 {
		return true
	}

	return remark[len(remark)-1:len(remark)] == "?"
}
