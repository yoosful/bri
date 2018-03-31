// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"crypto/sha256"
)

// User represents a registered user
type User struct {
	// Encrypted user address
	Address []byte `json:"address"`

	// User full name
	Name string `json:"name"`

	// User phone number
	Phone string `json:"phone"`

	// User usage summary
	Usage map[string]int `json:"usage"`
}

type Users struct {
	Data []User
}

// EncryptUser encrypts user info into hash using SHA256
func EncryptUser(name string, phone string) ([]byte, error) {
	h := sha256.New()
	src := make([]byte, 0, 256)

	if name == "" || phone == "" {
		return nil, ErrInvalidArguments
	}

	src = append(src, name...)
	src = append(src, phone...)

	if _, err := h.Write(src); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
