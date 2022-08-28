package main

import (
	"io/ioutil"
	"log"
)

// ReadFile returns content of a file.
func ReadFile(infile string) (output string) {
	content, err := ioutil.ReadFile(infile)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
