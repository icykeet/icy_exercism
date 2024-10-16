// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// This function validate input's name and return formatted output.
func ShareWith(name string) string {
	prf := fmt.Sprintf
	forMe := "one for me"
	var currentName string
	switch name {
	case "":
		currentName = "you"
	default:
		currentName = name
	}
	return prf("One for %s, %s.", currentName, forMe)
}
