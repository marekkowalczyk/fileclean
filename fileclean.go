package main

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {

	// fmt.Println(len(os.Args))
	fmt.Println(normalizeUTF(cleanText(removeIllFormed(readfile()))))
	// fmt.Println(cleanText(removeIllFormed(strings.Join(os.Args[1:], " "))))
	// fmt.Printf("%q\n", readfile())
	// fmt.Println(len(cleanText(readfile())))
	//writefile(cleanfile(readfile()))
}

func checkArguments() {
	if len(os.Args) == 1 {
		log.Fatal("No filename given; aborting")
	} else if len(os.Args) > 2 {
		log.Fatal("Too many arguments; aborting")
	}
}

func readfile() string {
	checkArguments()
	// infile := "test.txt"
	infile := os.Args[1]

	content, err := ioutil.ReadFile(infile)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func removeIllFormed(input string) (output string) {
	output, _, _ = transform.String(runes.ReplaceIllFormed(), input)
	return output
}

func cleanText(input string) (output string) {

	//fmt.Println(len(invisibleChars))

	clean := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, input)

	return clean

}

func normalizeUTF(input string) (output string) {
	normalizer := transform.Chain(norm.NFC)
	// normalizer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, _ := transform.String(normalizer, input)
	return normalized
}

func writefile(contentString string) {

	// Get byte data to write to file.
	dataString := contentString
	dataBytes := []byte(dataString)

	// Use WriteFile to create a file with byte data.
	err := ioutil.WriteFile("example0.txt", dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")

}
