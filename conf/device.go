// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"crypto/sha256"
	"encoding/hex"
)

// Device represents a registered device
type Device struct {
	// Encrypted address
	Address string `json: "address"`

	// Device type
	Dtype string `json: "dtype"`

	// Device id
	Did string `json: "did"`

	// Divece status
	Status bool `json: "status"`

	// Device payment rate
	Rate float64 `json:"rate"`

	// Last accessed user
	User string `,json:"user"`

	// Required privilege of the device
	// 0 - every user can access
	// 1 - only allowed users can access
	Privilege int

	// Addresses of permitted users
	Perm map[string]string `,json:"perm"`
}

type Devices struct {
	Data []Device
}

type deviceShort struct {
	dtype string
	did   string
	rate  float64
}

type DeviceMsg struct {
	Dtype string   `json:"dtype"`
	Did   string   `json:"did"`
	UInfo []string `json:"uid"`
	Msg   string   `json:"msg"`
}

// EncryptDevice encrypts a device using SHA256
func EncryptDevice(dtype, did string) (string, error) {
	h := sha256.New()
	src := make([]byte, 0, 256)

	if dtype == "" || did == "" {
		return "", ErrInvalidArguments
	}

	src = append(src, dtype...)
	src = append(src, did...)

	if _, err := h.Write(src); err != nil {
		return "", err
	}

	address := hex.EncodeToString(h.Sum(nil))

	return address, nil
}
