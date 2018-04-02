// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
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
	Usage map[string]float64 `json:"usage"`

	// Accessible devices
	Accessible []string `json:"accessible"`
}

type Users struct {
	Data []User
}

// UserStatusEntry represents information of a device used by a user
type UserStatusEntry struct {
	Dtype string  `json:"dtype"`
	Did   string  `json:"did"`
	Rate  float64 `json:"rate"`
	Used  float64 `json:"used"`
}

// UserStatus is sent as a response to user
type UserStatus struct {
	UserStatusList []UserStatusEntry `json:"list"`
	Price          float64           `json:"price"`
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

// GetTotalPrice will return the status of user
func (u User) GetStatus(address []byte) (UserStatus, error) {
	var user *User = nil

	for _, u := range UserData.Data {
		if bytes.Equal(user.Address, address) {
			user = &u
		}
	}

	// no matching user
	if user == nil {
		return UserStatus{}, ErrNoMatchingUser
	}

	var statlist []UserStatusEntry
	var price float64 = 0.0

	for dAddress, used := range user.Usage {
		h, err := hex.DecodeString(dAddress)
		device, err := DeviceData.Find(h)

		if err != nil {
			return UserStatus{}, err
		}

		// append usage information of a device
		statlist = append(statlist, UserStatusEntry{
			Dtype: device.dtype,
			Did:   device.did,
			Rate:  device.rate,
			Used:  used,
		})

		// accumulate user fee
		price += used * device.rate
	}

	return UserStatus{
		UserStatusList: statlist,
		Price:          price,
	}, nil
}
