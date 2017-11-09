package dev

import (
	"github.com/geoin/ble"
	"github.com/geoin/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
