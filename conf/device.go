// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"crypto/sha256"
)

// Device represents a registered device
type Device struct {
	// Encrypted address
	Address []byte `json: "address"`

	// Device type
	Dtype string `json: "dtype"`

	// Device id
	Did int64 `json: "did"`

	// Divece status
	Status bool `json: "status"`

	// Device payment rate
	Rate int `json:"rate"`

	// Last accessed user
	User []byte `,json:"user"`

	// Required privillege of the device
	// 0 - every user can access
	// 1 - only allowed users can access
	Privillege int

	// Addresses of permitted users
	Perm []string `,json:"perm"`
}

type Devices struct {
	Data []Device
}

// EncryptDevice encrypts a device using SHA256
func EncryptDevice(dtype, did string) ([]byte, error) {
	h := sha256.New()
	src := make([]byte, 0, 256)

	if dtype == "" || did == "" {
		return nil, ErrInvalidArguments
	}

	src = append(src, dtype...)
	src = append(src, did...)

	if _, err := h.Write(src); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
