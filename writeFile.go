package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Not tested and not used at the moment.

func writeFile(contentString string) {

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
