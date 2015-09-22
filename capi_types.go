package leptonica

// #include <stdio.h>
// #include <stdlib.h>
// #include <leptonica/allheaders.h>
import "C"
import "sync"

type Pix struct {
	pix      *C.PIX
	disposed bool
	mutex    sync.Mutex
}

type Dimension struct {
	Wide int
	Height int
}
