package dev

import (
	"github.com/geoin/ble"
	"github.com/geoin/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
