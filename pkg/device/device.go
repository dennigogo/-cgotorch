package device

import (
	"github.com/dennigogo/cgotorch/internal/cgolibrary"
	"github.com/dennigogo/cgotorch/internal/device"
)

type Device interface {
	Device() *cgolibrary.TorchDevice
	TypeDevice() cgolibrary.TorchDeviceType
}

func New() (Device, error) {
	return device.New()
}
