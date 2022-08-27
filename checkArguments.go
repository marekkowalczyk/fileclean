package main

import (
	"log"
)

func checkArguments(cmdLineArgs []string) (filename string) {
	if len(cmdLineArgs) == 1 {
		log.Fatal("No filename given; aborting")
	} else if len(cmdLineArgs) > 2 {
		log.Fatal("Too many arguments; aborting")
	}
	return cmdLineArgs[1]

}
