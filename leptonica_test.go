package leptonica_test

import (
	lept "github.com/cosmo0920/leptonica-capi-go"
	"path/filepath"
	"testing"
)

func setUp() (*lept.Pix) {
	targetFile := filepath.Join("example", "伊号潜水艦.png")
	pix, _ := lept.PixRead(targetFile)
	return pix;
}

func TestVersion(t *testing.T) {
	result := lept.Version()
	if result == "" {
		t.Errorf("result = %v cannot empty string", result)
	}
}

func TestConvertRGBToGrayFast(t *testing.T) {
	pix := setUp()

	_, err := pix.ConvertRGBToGrayFast()

	if err != nil {
		t.Errorf("Could not convert specified pix to grayscale.")
	}
}

func TestConvertRGBToGrayMinMax(t *testing.T) {
	pix := setUp()

	_, err := pix.ConvertRGBToGrayMinMax(lept.L_CHOOSE_MAX)

	if err != nil {
		t.Errorf("Could not convert specified pix to grayscale.")
	}
}

func TestConvertRGBToGray(t *testing.T) {
	pix := setUp()

	_, err := pix.ConvertRGBToGray(0.35, 0.35, 0.3)

	if err != nil {
		t.Errorf("Could not convert specified pix to grayscale.")
	}
}
