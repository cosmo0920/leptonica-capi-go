package leptonica

// #include <stdio.h>
// #include <stdlib.h>
// #include <leptonica/allheaders.h>
import "C"

// BoxaGetBox :: Ptr Boxa -> int32 -> int32 -> Ptr C.BOX
func BoxaGetBox(t *C.BOXA, index int32, flag CopyFlag) *C.BOX {
	return C.boxaGetBox(t, C.l_int32(index), C.l_int32(flag))
}
