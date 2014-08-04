package main

import (
	lept "../leptonica"
	"fmt"
)

const abort = 3

func main() {
	version := lept.Version()
	fmt.Println("leptonica version: " + version)
}
