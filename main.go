package main

import (
	"flag"
	"fmt"
	"os"
)

var oldLib bool
var dummyBuf []byte
var etag = []byte{0xbe, 0xef, 0xba, 0xbe}

func main() {
	flag.BoolVar(&oldLib, "old", true, "Use old library")
	flag.Parse()

	var err error
	dummyBuf, err = os.ReadFile("source.bin")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if oldLib {
		oldServer()
		return
	}
	newServer()
}
