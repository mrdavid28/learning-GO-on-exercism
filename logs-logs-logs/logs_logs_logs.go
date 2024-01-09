package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '*':
			return "whether"
		}

	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string { //three arguments "log" of type string also oldRune and NewRune of type rune
	oldRune_to_string := string(oldRune)
	newRune_to_string := string(newRune)
	return strings.Replace(log, oldRune_to_string, newRune_to_string, -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var log_length int = utf8.RuneCountInString(log)
	if log_length > limit {
		return false
	}
	return true
}

func main() {
	fmt.Println(WithinLimit("exercism❗", 9))
}
