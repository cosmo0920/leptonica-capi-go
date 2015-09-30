package leptonica_test

import (
	lept "github.com/cosmo0920/leptonica-capi-go"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func setUp() *lept.Pix {
	targetFile := filepath.Join("_example", "伊号潜水艦.png")
	pix, _ := lept.PixRead(targetFile)
	return pix
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

func TestPixWrite(t *testing.T) {
	pix := setUp()

	tmpDir, _ := ioutil.TempDir("", "temp-lept")
	tmpFilename := filepath.Join(tmpDir, "pixWrite.png")

	err := pix.PixWrite(tmpFilename, lept.IFF_PNG)

	if err != nil {
		t.Errorf("Could not write specified pix contents.")
	}
}

func TestPixCopy(t *testing.T) {
	pix := setUp()

	cpix := pix.PixCopy()

	result := pix.PixEqual(cpix)

	if result != true {
		t.Errorf("Suspisious copy operation.")
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

func TestGetDimension(t *testing.T) {
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

func TestGetDepth(t *testing.T) {
	pix := setUp()

	depth := pix.GetDepth()

	if depth < 1 {
		t.Errorf("Suspisious getting depth operation. depth: %v", depth)
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
}

func TestTwoSidedEdgeFiterWithInvalidFilterOption(t *testing.T) {
	pix := setUp()

	tpix, err := pix.ConvertRGBToGrayFast()

	_, err = tpix.TwoSidedEdgeFilter(lept.L_ALL_EDGES)

	if err == nil {
		t.Errorf("Suspisious applying two sided filter.")
	}
}

func TestConvertTo1(t *testing.T) {
	pix := setUp()

	tpix, err := pix.ConvertTo1(0)

	if err != nil || tpix.RawPix() == nil {
		t.Errorf("Could not convert specified pix to 1bpp(binary).")
	}

	tpix, err = pix.ConvertTo1(220)

	if err != nil || tpix.RawPix() == nil {
		t.Errorf("Could not convert specified pix to 1bpp(binary).")
	}

	tpix, err = pix.ConvertTo1(256)

	if err != nil || tpix.RawPix() == nil {
		t.Errorf("Could not convert specified pix to 1bpp(binary).")
	}
}

func TestConvertTo1WithInvalidThreshold(t *testing.T) {
	pix := setUp()

	_, err := pix.ConvertTo1(300)

	if err == nil {
		t.Errorf("Suspisous convertTo1 operation. [What] %v", err)
	}
}

func TestConvertTo8(t *testing.T) {
	pix := setUp()

	tpix := pix.ConvertTo8(lept.HAS_COLOR_MAP)

	if tpix.RawPix() == nil {
		t.Errorf("Could not convert specified pix to 8bpp.")
	}

}

func TestConvertTo16(t *testing.T) {
	pix := setUp()

	pix8 := pix.ConvertTo8(lept.HAS_COLOR_MAP)

	tpix := pix8.ConvertTo16()

	if tpix.RawPix() == nil {
		t.Errorf("Could not convert specified pix to 16bpp.")
	}

}

func TestConvertTo16WithInvalidPix(t *testing.T) {
	pix := setUp()

	tpix := pix.ConvertTo16()

	if tpix.RawPix() != nil {
		t.Errorf("Suspisious ConvertTo16 operation detected.")
	}

}

func TestConvertTo32(t *testing.T) {
	pix := setUp()

	tpix := pix.ConvertTo32()

	if tpix.RawPix() == nil {
		t.Errorf("Could not convert specified pix to 32bpp.")
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

func TestScaleWithInvalid(t *testing.T) {
	pix := setUp()

	_, err := pix.Scale(-2.0, -2.0)

	if err == nil {
		t.Errorf("Suspicious scale operation. [What] %v", err)
	}
}

func TestBoxCreate(t *testing.T) {
	_, err := lept.BoxCreate(60, 60, 40, 20)

	if err != nil {
		t.Errorf("Could not create box.")
	}
}

func TestBoxaCreate(t *testing.T) {
	_, err := lept.BoxaCreate(3)

	if err != nil {
		t.Errorf("Could not create boxa(s).")
	}
}
