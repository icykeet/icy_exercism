package techpalace

import (
	s "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + s.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsLine := s.Repeat("*", numStarsPerLine)
	newLine := "\n"
	return starsLine + newLine + welcomeMsg + newLine + starsLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return s.TrimSpace(s.ReplaceAll(oldMsg, "*", ""))
}
