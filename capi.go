package leptonica

// #include <stdio.h>
// #include <stdlib.h>
// #include <leptonica/allheaders.h>
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

func Version() string {
	cVersion := C.getLeptonicaVersion()
	version := C.GoString(cVersion)
	return version
}

// PixRead :: FilePath -> Ptr Pix
func PixRead(filename string) (*Pix, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	cPix := C.pixRead(cFilename)

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pix := &Pix{pix: cPix}

	runtime.SetFinalizer(pix, (*Pix).finalize)
	return pix, nil
}

// private finalize function
func (t *Pix) finalize() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.disposed {
		C.pixDestroy(&t.pix)
		C.free(unsafe.Pointer(t.pix))
		t.disposed = true
	}
}

func (t *Pix) GetDimension() (*Dimension, error) {
	var w, h C.l_int32
	result := C.pixGetDimensions(t.pix, &w, &h, nil)
	dim := &Dimension{Wide: int(w), Height: int(h)}

	if result == C.TRUE {
		return nil, errors.New("cannot get demensions")
	}

	return dim, nil
}

func (t *Pix) RankFilterGray(h int, w int, rank float32) (*Pix, error) {
	cPix := C.pixRankFilterGray(t.pix,
		C.l_int32(h), C.l_int32(w), C.l_float32(rank))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixd := &Pix{pix: cPix}

	runtime.SetFinalizer(pixd, (*Pix).finalize)
	return pixd, nil
}

func (t *Pix) DilateGray(h int, w int) (*Pix, error) {
	cPix := C.pixDilateGray(t.pix,
		C.l_int32(h), C.l_int32(w))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ErodeGray(h int, w int) (*Pix, error) {
	cPix := C.pixErodeGray(t.pix,
		C.l_int32(h), C.l_int32(w))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ScaleGrayRank2(index int) (*Pix, error) {
	cPix := C.pixScaleGrayRank2(t.pix, C.l_int32(index))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ConvertRGBToLuminance() (*Pix, error) {
	cPix := C.pixConvertRGBToLuminance(t.pix)

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ConvertRGBToGray(red float32, green float32, blue float32) (*Pix, error) {
	cPix := C.pixConvertRGBToGray(
		t.pix,
		C.l_float32(red), C.l_float32(green), C.l_float32(blue))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ConvertRGBToGrayFast() (*Pix, error) {
	cPix := C.pixConvertRGBToGrayFast(t.pix)

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ConvertRGBToGrayMinMax(grayChoose GrayChoose) (*Pix, error) {
	cPix := C.pixConvertRGBToGrayMinMax(t.pix, C.l_int32(grayChoose))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ScaleGrayLinear(scalex float32, scaley float32) (*Pix, error) {
	cPix := C.pixScaleGrayLI(
		t.pix,
		C.l_float32(scalex), C.l_float32(scaley))

	if cPix == nil {
		return nil, errors.New("cannot scale gray linear *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}


func (t *Pix) ScaleColorLinear(scalex float32, scaley float32) (*Pix, error) {
	cPix := C.pixScaleColorLI(
		t.pix,
		C.l_float32(scalex), C.l_float32(scaley))

	if cPix == nil {
		return nil, errors.New("cannot scale gray linear *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ScaleGrayMinMax(xfact int, yfact int, grayChoose GrayChoose) (*Pix, error) {
	cPix := C.pixScaleGrayMinMax(
		t.pix,
		C.l_int32(xfact), C.l_int32(yfact), C.l_int32(grayChoose))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ScaleGrayMinMax2(grayChoose GrayChoose) (*Pix, error) {
	cPix := C.pixScaleGrayMinMax2(t.pix, C.l_int32(grayChoose))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ScaleGrayRankCascade(level1 int, level2 int, level3 int, level4 int) (*Pix, error) {
	cPix := C.pixScaleGrayRankCascade(
		t.pix,
		C.l_int32(level1), C.l_int32(level2),
		C.l_int32(level3), C.l_int32(level4))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

// Change Pix Scale with float32 values (x: float32, y: float32)
func (t *Pix) Scale(x float32, y float32) (*Pix, error) {
	cPix := C.pixScale(t.pix, C.l_float32(x), C.l_float32(y))

	if cPix == nil {
		return nil, errors.New("cannot create *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) RemoveColormap(colorMap ColorMap) (*Pix, error) {
	cPix := C.pixRemoveColormap(t.pix, C.l_int32(colorMap))

	if cPix == nil {
		return nil, errors.New("cannot remove color map from *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) ColorSegment(max_dist int, max_color int, sel_size int, final_colors int) (*Pix, error) {
	cPix := C.pixColorSegment(t.pix,
		C.l_int32(max_dist), C.l_int32(max_color),
		C.l_int32(sel_size), C.l_int32(final_colors))

	if cPix == nil {
		return nil, errors.New("cannot remove color map from *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) SobelEdgeFilter(orient OrientFlag) (*Pix, error) {
	cPix := C.pixSobelEdgeFilter(t.pix, C.l_int32(orient))

	if cPix == nil {
		return nil, errors.New("cannot apply sobel edge filter to *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}


func (t *Pix) TwoSidedEdgeFilter(orient OrientFlag) (*Pix, error) {
	cPix := C.pixTwoSidedEdgeFilter(t.pix, C.l_int32(orient))

	if cPix == nil {
		return nil, errors.New("cannot apply two sided edge filter to *Pix")
	}

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt, nil
}

func (t *Pix) PixEqual(dPix *Pix) (bool) {
	var same C.l_int32
	C.pixEqual(t.pix, dPix.pix, &same)

	if same == C.TRUE {
		return true
	}

	return false
}

func (t *Pix) AddBorder(npix int, color uint) (*Pix) {
	cPix := C.pixAddBorder(t.pix, C.l_int32(npix), C.l_uint32(color))

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt
}

func (t *Pix) RemoveBorder(npix int) (*Pix) {
	cPix := C.pixRemoveBorder(t.pix, C.l_int32(npix))

	pixt := &Pix{pix: cPix}

	runtime.SetFinalizer(pixt, (*Pix).finalize)
	return pixt
}

// PixWrite :: Ptr Pix -> String -> IMGFormat -> error
func (t *Pix) PixWrite(path string, format IMGFormat) (error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	result := C.pixWrite(cPath, t.pix, C.l_int32(format))

	if result == 0 {
		return errors.New("cannot write *Pix")
	}

	return nil
}

// RawPix :: Ptr Pix -> Ptr C.PIX
func (t *Pix) RawPix() *C.PIX {
	return t.pix
}

// BoxaGetBox :: Ptr Boxa -> int32 -> int32 -> Ptr C.BOX
func BoxaGetBox(t *C.BOXA, index int32, flag CopyFlag) *C.BOX {
	return C.boxaGetBox(t, C.l_int32(index), C.l_int32(flag))
}
