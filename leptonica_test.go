package leptonica_test

import (
	lept "github.com/cosmo0920/leptonica-capi-go"
	"path/filepath"
	"testing"
)

func setUp() (*lept.Pix) {
	targetFile := filepath.Join("_example", "伊号潜水艦.png")
	pix, _ := lept.PixRead(targetFile)
	return pix;
}

func TestVersion(t *testing.T) {
	result := lept.Version()
	if result == "" {
		t.Errorf("result = %v cannot empty string", result)
	}
}

func TestPixRead(t *testing.T) {
	targetFile := filepath.Join("_example", "伊号潜水艦.png")
	_, err := lept.PixRead(targetFile)

	if err != nil {
		t.Errorf("Could not read pix from specified img file.")
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

func TestConvertRGBToLuminance(t *testing.T) {
	pix := setUp()

	_, err := pix.ConvertRGBToLuminance()

	if err != nil {
		t.Errorf("Could not convert specified pix to grayscale.")
	}
}

func TestGetDimensions(t *testing.T) {
	pix := setUp()

	dim, err := pix.GetDimension()

	if err != nil {
		t.Errorf("Could not get dimensions from specified pix.")
	}

	if dim == nil {
		t.Errorf("Suspicious dimensions.")
	}
}

func TestAddBorder(t *testing.T) {
	pix := setUp()

	tpix := pix.AddBorder(5, 0)

	result := pix.PixEqual(tpix)

	if result == true {
		t.Errorf("Suspicious addBorder operation.")
	}
}

func TestRemoveBorder(t *testing.T) {
	pix := setUp()

	const BORDER_WIDTH = 5
	bpix := pix.AddBorder(BORDER_WIDTH, 0)

	tpix := bpix.RemoveBorder(BORDER_WIDTH)

	result := tpix.PixEqual(pix)

	if result != true {
		t.Errorf("Suspisious border operation.")
	}
}

func TestSobelEdgeFiter(t *testing.T) {
	pix := setUp()


	tpix, err := pix.ConvertRGBToGrayFast()

	if err != nil {
		t.Errorf("Could not convert specified pix to grayscale.")
	}

	_, err = tpix.SobelEdgeFilter(lept.L_ALL_EDGES)

	if err != nil {
		t.Errorf("Could not apply filter to specified pix.")
	}
}


func TestTwoSidedEdgeFiter(t *testing.T) {
	pix := setUp()


	tpix, err := pix.ConvertRGBToGrayFast()

	if err != nil {
		t.Errorf("Could not convert specified pix to grayscale.")
	}

	_, err = tpix.TwoSidedEdgeFilter(lept.L_VERTICAL_EDGES)

	if err != nil {
		t.Errorf("Could not apply filter to specified pix.")
	}

	_, err = tpix.TwoSidedEdgeFilter(lept.L_ALL_EDGES)

	if err == nil {
		t.Errorf("Suspisious applying two sided filter.")
	}
}

func TestRemoveColorMap(t *testing.T) {
	pix := setUp()

	_, err := pix.RemoveColormap(lept.REMOVE_CMAP_TO_GRAYSCALE)


	if err != nil {
		t.Errorf("Could not scale operation.")
	}
}

func TestScale(t *testing.T) {
	pix := setUp()

	_, err := pix.Scale(2.0, 2.0)


	if err != nil {
		t.Errorf("Could not scale operation.")
	}
}
