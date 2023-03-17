package cgolibrary

// #cgo CPPFLAGS: -I${SRCDIR}/libtorch/include
// #cgo CPPFLAGS: -I${SRCDIR}/libtorch/include/torch/csrc/api/include
// #cgo LDFLAGS: -L${SRCDIR}/libtorch/lib -Wl,-rpath ${SRCDIR}/libtorch/lib -lc10 -ltorch -ltorch_cpu
// #include "cgolibrary.h"
import "C"

import (
	"unsafe"

	"github.com/pkg/errors"
)

// ErrUnsafePointerInvalid error: unsafe pointer must be valid.
var ErrUnsafePointerInvalid = errors.New("unsafe pointer must be valid")

// E returns error.
func E(err unsafe.Pointer) error {
	if err != nil {
		msg := C.GoString((*C.char)(err))
		defer C.FreeString((*C.char)(err))
		return errors.Wrap(ErrUnsafePointerInvalid, msg)
	}

	return nil
}
