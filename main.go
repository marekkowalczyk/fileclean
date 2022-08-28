package main

import (
	"fmt"
	"os"
)

// The fileclean command prints to STDOUT cleaned-up text from a file given as an argument. To change file in place use 'fileclean file.txt | sponge file.txt'.
func main() {
	fmt.Println(NormalizeUTF(CleanText(RemoveIllFormed(ReadFile(GetFilename(os.Args))))))
}
