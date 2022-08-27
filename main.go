package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(normalizeUTF(cleanText(removeIllFormed(readFile(checkArguments(os.Args))))))
}
