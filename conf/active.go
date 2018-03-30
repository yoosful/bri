// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"bytes"
	"time"
)

var Actives []Active

// Record contains records of active devices
type Active struct {
	Address []byte
	On      time.Time
}

func GetDuration(actives []Active, address []byte) (time.Duration, error) {
	for _, active := range actives {
		if bytes.Equal(address, active.Address) {
			return time.Since(active.On), nil
		}
	}

	return time.Duration(0), ErrNoMathingDevice
}
