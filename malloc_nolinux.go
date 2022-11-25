//go:build !linux
// +build !linux

package malloc

// #cgo CFLAGS: -std=c11
// #include "stdlib.h"
import "C"
import "unsafe"

type MalloptParam int

const (
	ArenaMax      MalloptParam = 0
	ArenaTest     MalloptParam = 0
	CheckAction   MalloptParam = 0
	MMapMax       MalloptParam = 0
	MMapThreshold MalloptParam = 0
	MXFast        MalloptParam = 0
	Perturb       MalloptParam = 0
	TopPad        MalloptParam = 0
	TrimThreshold MalloptParam = 0
)

func MallocTrim(pad int) bool {
	return true
}

func Mallopt(param MalloptParam, value int) bool {
	return true
}

func Malloc(size int) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}
