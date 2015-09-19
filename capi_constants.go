package leptonica

// #include <stdio.h>
// #include <stdlib.h>
// #include <leptonica/allheaders.h>
import "C"

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

type GrayChoose int32

const (
	L_CHOOSE_MIN = 1 + iota
	L_CHOOSE_MAX
	L_CHOOSE_MAX_MIN_DIFF
)

type ColorMap int32

const (
	REMOVE_CMAP_TO_BINARY = iota
	REMOVE_CMAP_TO_GRAYSCALE
	REMOVE_CMAP_TO_FULL_COLOR
	REMOVE_CMAP_BASED_ON_SRC
)

type OrientFlag int32

const (
	L_HORIZONTAL_EDGES = iota
	L_VERTICAL_EDGES
	L_ALL_EDGES
)
