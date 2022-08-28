package main

import (
	"log"
)

// GetFilename checks the array of command line arguments for correct number of input files (one) and returns the filename as string.
func GetFilename(cmdLineArgs []string) (filename string) {
	if len(cmdLineArgs) == 1 {
		log.Fatal("No filename given; aborting")
	} else if len(cmdLineArgs) > 2 {
		log.Fatal("Too many arguments; aborting")
	}
	return cmdLineArgs[1]

}
