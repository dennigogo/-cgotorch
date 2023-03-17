package cpu

import (
	"github.com/dennigogo/cgotorch/internal/cgolibrary"
	"github.com/pkg/errors"
)

type Cpu interface {
	Device() *cgolibrary.TorchDevice
	TypeDevice() cgolibrary.TorchDeviceType
}

type cpu struct {
	D *cgolibrary.TorchDevice
}

func New() (Cpu, error) {
	d, err := cgolibrary.NewDevice(cgolibrary.CpuDevice)
	if err != nil {
		const em = "new cpu device"
		return nil, errors.Wrap(err, em)
	}

	return &cpu{D: d.Device()}, nil
}

func (c *cpu) Device() *cgolibrary.TorchDevice {
	return c.D
}

func (c *cpu) TypeDevice() cgolibrary.TorchDeviceType {
	return cgolibrary.CpuDevice
}
