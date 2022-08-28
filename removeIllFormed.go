package main

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

// RemoveIllFormed removes ill-formed UTF runes from input text.
func RemoveIllFormed(input string) (output string) {
	output, _, _ = transform.String(runes.ReplaceIllFormed(), input)
	return output
}
