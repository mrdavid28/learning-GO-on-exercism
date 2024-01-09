package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	regexk := regexp.MustCompile(text)
	return regexk.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
}

func SplitLogLine(text string) []string {
	panic("Please implement the SplitLogLine function")
}

func CountQuotedPasswords(lines []string) int {
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
