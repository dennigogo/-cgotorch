package cgolibrary

// #cgo CPPFLAGS: -I${SRCDIR}/libtorch/include
// #cgo CPPFLAGS: -I${SRCDIR}/libtorch/include/torch/csrc/api/include
// #cgo LDFLAGS: -L${SRCDIR}/libtorch/lib -Wl,-rpath ${SRCDIR}/libtorch/lib -lc10 -ltorch -ltorch_cpu
// #include "cgolibrary.h"
import "C"

func IsCUDAAvailable() bool {
	return C.IsCUDAAvailable() == C.bool(true)
}
