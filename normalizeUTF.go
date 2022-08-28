package main

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//NormalizeUTF converts input string to output string by substituting all instances of NFD Unicode character n-tuples with canonically equivalent NFC forms, where available.
func NormalizeUTF(input string) (output string) {
	normalizer := transform.Chain(norm.NFC)
	// normalizer := transform.Chain(norm.NFC)
	// normalizer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, _ := transform.String(normalizer, input)
	return normalized
}
