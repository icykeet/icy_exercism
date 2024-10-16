package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[[TRCDBGINFWREFL]{3}\]).*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<[*~\-=]*\>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`\"(.)*[paswordPASWORD]{8}\"`)
	counter := 0
	for _, line := range lines {
		if re.MatchString(line) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line)(\d+)`)
	cleaned := re.ReplaceAllString(text, "")
	return cleaned
}

func TagWithUserName(lines []string) []string {
	linesN := []string{}
	for _, message := range lines {
		reMust := regexp.MustCompile(`User +\w*`)

		isValidMessage := reMust.MatchString(message)
		ValidMessage := reMust.FindString(message)

		if isValidMessage {
			reFS := regexp.MustCompile(`^User +(\w+)`)
			login := reFS.FindStringSubmatch(ValidMessage)

			if login != nil {
				userTag := "[USR] " + login[1] + " " + message
				linesN = append(linesN, userTag)
			}
		} else {
			linesN = append(linesN, message)
		}
	}
	return linesN
}
