package main

import (
	"io/ioutil"
	"log"
)

func readFile(infile string) (output string) {
	content, err := ioutil.ReadFile(infile)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
