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

type TorchDeviceType string

const (
	CpuDevice  TorchDeviceType = "cpu"
	CudaDevice TorchDeviceType = "cuda"
)

// Device describes of device.
type Device interface {
	Device() *TorchDevice
}

type TorchDevice struct {
	CD C.Device
}

// Device wrappers a pointer to C.Device
type device struct {
	TD *TorchDevice
}

// NewDevice returns a Device.
func NewDevice(deviceType TorchDeviceType) (Device, error) {
	var cd C.Device
	if err := E(unsafe.Pointer(C.Torch_Device(C.CString(string(deviceType)), &cd))); err != nil {
		const em = "new device"
		return nil, errors.Wrap(err, em)
	}

	return &device{TD: &TorchDevice{CD: cd}}, nil
}

// Device returns device.
func (d *device) Device() *TorchDevice {
	return d.TD
}
