package leptonica
// #include <leptonica/allheaders.h>
//
// l_int32 getNumberOfBox(BOXA* boxa) {
// 	return boxa->n;
// }
import "C"

func GetNumberOfBox(t *C.BOXA) int {
	return int(C.getNumberOfBox(t))
}
