package conf

import (
	"crypto/sha256"
	"errors"
)

var ErrInvalidArguments = errors.New("Invalid arguments")

type Device struct {
	Address []byte         `json: "address"`
	dtype   string         `json: "dtype"`
	did     int64          `json: "did"`
	usage   map[string]int `json: "usage"`
}

func Checksum(dtype, did string) ([]byte, error) {
	h := sha256.New()
	src := make([]byte, 256)

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
