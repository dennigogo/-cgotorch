package cuda

import (
	"github.com/dennigogo/cgotorch/internal/cgolibrary"
	"github.com/pkg/errors"
)

type Cuda interface {
	Device() *cgolibrary.TorchDevice
	TypeDevice() cgolibrary.TorchDeviceType
}

type cuda struct {
	D *cgolibrary.TorchDevice
}

func New() (Cuda, error) {
	c := cuda{}

	if cgolibrary.IsCUDAAvailable() {
		d, err := cgolibrary.NewDevice(cgolibrary.CudaDevice)
		if err != nil {
			const em = "new cuda device"
			return nil, errors.Wrap(err, em)
		}

		c.D = d.Device()
	}

	return &c, nil
}

func (c *cuda) Device() *cgolibrary.TorchDevice {
	return c.D
}

func (c *cuda) TypeDevice() cgolibrary.TorchDeviceType {
	return cgolibrary.CudaDevice
}
