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

	fmt.Println("ConvertTo8:", targetFile)
	tpix := pix.ConvertTo8(lept.HAS_COLOR_MAP)

	if err != nil {
		panic("Could not convert specified pix to grayscale.")
	}

	fmt.Println("Apply SobelEdgeFilter:", targetFile)
	epix, err := tpix.SobelEdgeFilter(lept.L_ALL_EDGES)
	epix.PixWrite("SobelEdge.png", lept.IFF_PNG)
}
