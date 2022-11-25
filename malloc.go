//go:build linux
// +build linux

package malloc

// #cgo CFLAGS: -std=c11
// #include "malloc.h"
import "C"
import "unsafe"

type MalloptParam int

const (
	ArenaMax      MalloptParam = C.M_ARENA_MAX
	ArenaTest     MalloptParam = C.M_ARENA_TEST
	CheckAction   MalloptParam = C.M_CHECK_ACTION
	MMapMax       MalloptParam = C.M_MMAP_MAX
	MMapThreshold MalloptParam = C.M_MMAP_THRESHOLD
	MXFast        MalloptParam = C.M_MXFAST
	Perturb       MalloptParam = C.M_PERTURB
	TopPad        MalloptParam = C.M_TOP_PAD
	TrimThreshold MalloptParam = C.M_TRIM_THRESHOLD
)

func MallocTrim(pad int) bool {
	return C.malloc_trim(C.size_t(pad)) > 0
}

func Mallopt(param MalloptParam, value int) bool {
	return C.mallopt(C.int(param), C.int(value)) > 0
}

func Malloc(size int) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}
