package device

import (
	"github.com/dennigogo/cgotorch/internal/cgolibrary"
	"github.com/dennigogo/cgotorch/internal/device/cpu"
	"github.com/dennigogo/cgotorch/internal/device/cuda"
	"github.com/pkg/errors"
)

type Device interface {
	Device() *cgolibrary.TorchDevice
	TypeDevice() cgolibrary.TorchDeviceType
}

type device struct {
	TD *cgolibrary.TorchDevice
}

func New() (Device, error) {
	d := device{}
	td, err := d.recognize()
	if err != nil {
		const em = "recognize device"
		return nil, errors.Wrap(err, em)
	}

	return td, nil
}

func (d *device) recognize() (Device, error) {
	switch {
	case cgolibrary.IsCUDAAvailable():
		td, err := cuda.New()
		if err != nil {
			const em = "new cuda device (recognize)"
			return nil, errors.Wrap(err, em)
		}

		return td, nil
	default:
		td, err := cpu.New()
		if err != nil {
			const em = "new cpu device (recognize)"
			return nil, errors.Wrap(err, em)
		}

		return td, nil
	}
}
