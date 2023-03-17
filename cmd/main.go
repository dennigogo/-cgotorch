package main

import (
	"log"

	"github.com/dennigogo/cgotorch/pkg/device"
)

func main() {
	d, err := device.New()
	if err != nil {
		panic(err)
	}

	log.Printf("device %+v | type = %s", d.Device(), d.TypeDevice())
}
