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
		fmt.Println("Could not read specified png file.")
	}

	fmt.Println("ConvertRGBToGrayFast:", targetFile)
	tpix, err := pix.ConvertRGBToGrayFast()

	if err != nil {
		fmt.Println("Could not convert specified pix to grayscale. Panic.")
	}

	tpix.PixWrite("Translated.png", lept.IFF_PNG)
}
