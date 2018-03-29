// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"crypto/sha256"
	"errors"
)

var ErrInvalidArguments = errors.New("Invalid arguments")

type Device struct {
	Address []byte         `json: "address"`
	Dtype   string         `json: "dtype"`
	Did     int64          `json: "did"`
	Status  bool           `json: "status"`
	Usage   map[string]int `,json: "usage"`
}

type Devices struct {
	Data []Device
}

func Checksum(dtype, did string) ([]byte, error) {
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
