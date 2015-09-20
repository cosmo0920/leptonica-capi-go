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

	fmt.Println("ConvertRGBToGrayFast:", targetFile)
	tpix, err := pix.ConvertRGBToGrayFast()

	if err != nil {
		panic("Could not convert specified pix to grayscale.")
	}

	result := pix.PixEqual(tpix)

	if result == true {
		panic("Suspicious pix conversion.")
	}

	tpix.PixWrite("Translated.png", lept.IFF_PNG)
}
