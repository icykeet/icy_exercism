package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		switch v {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog []rune
	for _, value := range log {
		if value == oldRune {
			newLog = append(newLog, newRune)
		} else {
			newLog = append(newLog, value)
		}
	}
	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	actualLength := 0
	for range log {
		actualLength += 1
	}
	return actualLength <= limit
	// return utf8.RuneCountInString(log) <= limit -- solution in one line
}
