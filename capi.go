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

type CopyFlag int32

const (
	L_INSERT CopyFlag = iota
	L_COPY
	L_CLONE
	L_COPY_CLONE
)

type IMGFormat int32
const (
	IFF_UNKNOWN IMGFormat = iota
	IFF_BMP
	IFF_JFIF_JPEG
	IFF_PNG
	IFF_TIFF
	IFF_TIFF_PACKBITS
	IFF_TIFF_RLE
	IFF_TIFF_G3
	IFF_TIFF_G4
	IFF_TIFF_LZW
	IFF_TIFF_ZIP
	IFF_PNM
	IFF_PS
	IFF_GIF
	IFF_JP2
	IFF_WEBP
	IFF_LPDF
	IFF_DEFAULT
	IFF_SPIX
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

func (t *Pix) finalize() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.disposed {
		C.pixDestroy(&t.pix)
		C.free(unsafe.Pointer(t.pix))
		t.disposed = true
	}
}

func (t *Pix) GetDimension() (int, int) {
	var w, h C.l_int32
	C.pixGetDimensions(t.pix, &w, &h, nil)
	return int(w), int(h)
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

func (t *Pix) PixWrite(path string, format IMGFormat) (error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	result := C.pixWrite(cPath, t.pix, C.l_int32(format))

	if result == 0 {
		return errors.New("cannot write *Pix")
	}

	return nil
}

// rawPix :: Ptr Pix -> Ptr C.PIX
func (t *Pix) RawPix() *C.PIX {
	return t.pix
}

// BoxaGetBox :: Ptr Boxa -> int32 -> int32 -> Ptr C.BOX
func BoxaGetBox(t *C.BOXA, index int32, flag CopyFlag) *C.BOX {
	return C.boxaGetBox(t, C.l_int32(index), C.l_int32(flag))
}
