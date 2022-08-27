package main

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func normalizeUTF(input string) (output string) {
	normalizer := transform.Chain(norm.NFC)
	// normalizer := transform.Chain(norm.NFC)
	// normalizer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, _ := transform.String(normalizer, input)
	return normalized
}
