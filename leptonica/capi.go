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

func (t *Pix) finalize() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.disposed {
		C.pixDestroy(&t.pix)
		C.free(unsafe.Pointer(t.pix))
		t.disposed = true
	}
}

// rawPix :: Ptr Pix -> Ptr C.PIX
func (t *Pix) RawPix() *C.PIX {
	return t.pix
}
