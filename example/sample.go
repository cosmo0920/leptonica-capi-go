package main

import (
	lept "github.com/cosmo0920/leptonica-capi-go"
	"fmt"
)

func main() {
	version := lept.Version()
	fmt.Println("leptonica version: " + version)
}
