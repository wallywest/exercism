package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func Hey(input string) string {
	i := strings.TrimSpace(input)

	switch {
	case isSilent(i):
		return "Fine. Be that way!"
	case isUppercase(i) && !isLowerCase(i):
		return "Whoa, chill out!"
	case isQuestion(i):
		return "Sure."
	default:
		return "Whatever."
	}
}

func isQuestion(input string) bool {
	if len(input) == 0 {
		return false
	}
	return input[len(input)-1] == '?'
}

func isUppercase(input string) bool {
	for _, i := range input {
		if unicode.IsUpper(i) {
			return true
		}
	}

	return false
}

func isLowerCase(input string) bool {
	for _, i := range input {
		if unicode.IsLower(i) {
			return true
		}
	}

	return false
}

func isSilent(input string) bool {
	if input == "" {
		return true
	}
	return false
}
