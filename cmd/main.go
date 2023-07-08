package main

import (
	"flag"
	"log"

	"pastecode-cli/internal"
)

func main() {
	log.SetFlags(0)
	fileFlag := flag.String("f", "", "File path to read")
	flag.Parse()

	contents, err := internal.GetContents(fileFlag)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(contents))
}
