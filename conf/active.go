// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"bytes"
	"time"
)

// Actives holds runtime information of activated devices
var Actives []Active

// Record contains records of active devices
type Active struct {
	Address []byte
	On      time.Time
}

func GetDuration(actives []Active, address []byte) (time.Duration, error) {
	i := 0
	for _, active := range actives {
		if bytes.Equal(address, active.Address) {
			break
		}
		i++
	}

	if i < len(Actives) {
		d := time.Since(Actives[i].On)
		Actives = append(Actives[:i], Actives[i+1:]...)
		return d, nil
	}

	// currently if there is no matching device, GetDuration
	// will transparently return time.Duration(0)
	return time.Duration(0), nil
}

func SetOnTime(actives []Active, address []byte) error {
	for _, active := range actives {
		if bytes.Equal(address, active.Address) {
			return ErrUnexpectedBehavior
		}
	}

	Actives = append(Actives, Active{
		Address: address,
		On:      time.Now(),
	})

	return nil
}
