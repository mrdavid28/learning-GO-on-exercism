package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return string("Welcome to the Tech Palace, " + strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return string(strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine))
}

// CleanupMessage cleans up an old marketing message.
/*
*There is a function to [replace strings within a `string`][replace-all].
- There is a function to [trim leading and trailing spaces from a `string`][trim-space].
*/
func CleanupMessage(oldMsg string) string {
	newMsg := strings.ReplaceAll(oldMsg, "*", "")
	newMsg = strings.TrimSpace(newMsg)
	return newMsg
}
