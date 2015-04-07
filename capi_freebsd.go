// +build freebsd

package leptonica

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -llept
import "C"

/*
Note: specify LDFLAGS by hand for workaround.
BUG: FreeBSD 10 release do not provide leptonica with lept.pc.
*/
