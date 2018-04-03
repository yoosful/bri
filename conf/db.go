// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	DeviceData Devices
	UserData   Users
)

func (d *Devices) Init() error {
	file, err := os.OpenFile("bri-devices.json", os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(raw, &d.Data)
	return nil
}

func (d *Devices) EncryptAndAdd(did string, dtype string, status bool) error {
	h, err := EncryptDevice(dtype, did)
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range d.Data {
		if h == device.Address {
			return ErrDuplicateDevice
		}
	}

	d.Data = append(d.Data, Device{
		Address:   h,
		Dtype:     dtype,
		Did:       did,
		Status:    status,
		Privilege: 0,
		Rate:      float64(1.0),
	})
	return nil
}

func (d *Devices) Add(address string, new Device) error {
	for _, device := range d.Data {
		if address == device.Address {
			return ErrDuplicateDevice
		}
	}

	d.Data = append(d.Data, new)
	return nil
}

func (d *Devices) Find(address string) (deviceShort, error) {
	for _, device := range DeviceData.Data {
		if address == device.Address {
			return deviceShort{
				dtype: device.Dtype,
				did:   device.Did,
				rate:  device.Rate,
			}, nil
		}
	}

	return deviceShort{}, ErrNoMatchingDevice
}

func (d *Devices) UpdateStatus(address string, user string, msg string) error {
	i := 0
	for _, device := range d.Data {
		if address == device.Address {
			isTurnedOn := d.Data[i].Status

			if isTurnedOn {
				// Turning off the device
				if msg == "on" {
					log.Println("Already Turned On")
					break
				} else if msg != "off" {
					log.Println("Unexpected Message Arrived")
					break
				}
				if device.User != user {
					log.Println("Turned Off by a Different User?!")
					break
				} else {
					d.Data[i].Status = false

					// get active duration to update user's usage info
					duration, err := GetDuration(Actives, address)
					if err != nil {
						panic(err)
					}

					UserData.UpdateUsage(user, address, duration.Seconds())
				}
			} else {
				// Turning on the device
				if msg == "off" {
					log.Println("Not Turned On Yet")
					break
				} else if msg != "on" {
					log.Println("Unexpected Message Arrived")
					break
				} else {
					d.Data[i].User = user
					d.Data[i].Status = true

					// set device on time
					err := SetOnTime(Actives, address)
					if err != nil {
						panic(err)
					}
				}
			}
			break
		}
		i++
	}

	if i >= len(d.Data) {
		log.Println("No matching address")
	}

	return nil
}

func (d *Devices) Delete(address string) error {
	i := 0
	for _, device := range d.Data {
		if address == device.Address {
			break
		}
		i++
	}

	if i < len(d.Data) {
		d.Data = append(d.Data[:i], d.Data[i+1:]...)
	} else {
		return ErrNoMatchingDevice
	}

	return nil
}

func (d *Devices) Dump() error {
	bytes, err := json.Marshal(&d.Data)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("bri-devices.json", bytes, 0666)
	return nil
}

func (u *Users) Init() error {
	file, err := os.OpenFile("bri-users.json", os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(raw, &u.Data)

	return nil
}

func (u *Users) EncryptAndAdd(name, phone string) error {
	h, err := EncryptUser(name, phone)
	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range u.Data {
		if h == user.Address {
			return ErrDuplicateUser
		}
	}

	u.Data = append(u.Data, User{
		Address:     h,
		Name:        name,
		Phone:       phone,
		Usage:       map[string]float64{},
		Priviledged: map[string]string{},
	})

	return nil
}

func (u *Users) Delete(address string) error {
	i := 0
	for _, user := range u.Data {
		if address == user.Address {
			break
		}
		i++
	}

	if i < len(u.Data) {
		u.Data = append(u.Data[:i], u.Data[i+1:]...)
	} else {
		return ErrNoMatchingUser
	}

	return nil
}

// UpdateUsage updage usage info of a user with device id and
// the amount of time turned on.
func (u *Users) UpdateUsage(address string, device string, amount float64) error {
	i := 0
	for _, user := range u.Data {
		if address == user.Address {
			break
		}
		i++
	}

	if i >= len(u.Data) {
		return ErrNoMatchingUser
	}

	j := 0
	for k, _ := range u.Data[i].Usage {
		if k == device {
			u.Data[i].Usage[k] += amount
			break
		}
		j++
	}

	if j >= len(u.Data[i].Usage) {
		u.Data[i].Usage[device] = amount
	}

	return nil
}

func (u *Users) Dump() error {
	bytes, err := json.Marshal(&u.Data)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("bri-users.json", bytes, 0666)
	return nil
}
