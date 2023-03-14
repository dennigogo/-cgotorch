package cgolibrary

// #cgo CPPFLAGS: -I${SRCDIR}/libtorch/include
// #cgo CPPFLAGS: -I${SRCDIR}/libtorch/include/torch/csrc/api/include
// #cgo LDFLAGS: -L${SRCDIR}/libtorch/lib -Wl,-rpath ${SRCDIR}/libtorch/lib -lc10 -ltorch -ltorch_cpu
// #include "cgolibrary.h"
import "C"

import (
	"fmt"
)

func Test() {
	if C.IsCUDAAvailable() == C.bool(true) {
		fmt.Printf("Is CUDA\n")
		return
	}

	fmt.Printf("Is Not CUDA\n")
}
