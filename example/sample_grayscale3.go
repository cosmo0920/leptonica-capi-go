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

	fmt.Println("ConvertRGBToGray:", targetFile)
	tpix, err := pix.ConvertRGBToGray(0.35, 0.35, 0.3)

	if err != nil {
		panic("Could not convert specified pix to grayscale.")
	}

	result := pix.PixEqual(tpix)

	if result == true {
		panic("Suspicious pix conversion.")
	}

	tpix.PixWrite("Translated3.png", lept.IFF_PNG)
}
