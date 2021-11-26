package main

import "strings"

const (
	OPEN_BRACKETS  = "([{"
	CLOSE_BRACKETS = ")]}"
)

func isOpenBracket(c rune) bool {
	return strings.ContainsRune(OPEN_BRACKETS, c)
}

func isCloseBracket(c rune) bool {
	return strings.ContainsRune(CLOSE_BRACKETS, c)
}

func getOpenBracket(c rune) rune {
	i := strings.IndexRune(CLOSE_BRACKETS, c)
	return rune(OPEN_BRACKETS[i])
}

func IsBalanced(input string) bool {
	var stack []rune

	for _, c := range input {

		if isOpenBracket(c) {
			stack = append(stack, c)
		} else if isCloseBracket(c) {
			l := len(stack)

			if l == 0 || stack[l-1] != getOpenBracket(c) {
				return false
			}
			stack = stack[:l-1]
		}
	}

	return len(stack) == 0
}
