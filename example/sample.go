package main

import (
	lept "../leptonica"
	"fmt"
)

func main() {
	version := lept.Version()
	fmt.Println("leptonica version: " + version)
}
