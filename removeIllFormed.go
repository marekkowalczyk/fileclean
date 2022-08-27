package main

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

func removeIllFormed(input string) (output string) {
	output, _, _ = transform.String(runes.ReplaceIllFormed(), input)
	return output
}
