package main

import (
	"strings"
	"unicode"
)

// CleanText removes non-printing and non-space characters from input.
func CleanText(input string) (output string) {

	//fmt.Println(len(invisibleChars))

	clean := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, input)

	return clean

}
