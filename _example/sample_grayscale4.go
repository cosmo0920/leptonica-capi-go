package main

import (
	"fmt"
	lept "github.com/cosmo0920/leptonica-capi-go"
)

func main() {
	version := lept.Version()
	fmt.Println("leptonica version: " + version)
	targetFile := "伊号潜水艦.png"
	pix, err := lept.PixRead(targetFile)

	if err != nil {
		panic("Could not read specified png file.")
	}

	fmt.Println("ConvertRGBToLuminance:", targetFile)
	tpix, err := pix.ConvertRGBToLuminance()

	if err != nil {
		panic("Could not convert specified pix to grayscale.")
	}

	tpix.PixWrite("Translated4.png", lept.IFF_PNG)
}
