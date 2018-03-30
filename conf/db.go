// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var (
	DeviceData Devices
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

	for _, device := range d.Data {
		if bytes.Equal(device.Address, h) {
			return ErrDuplicateDevice
		}
	}

	didInt, err := strconv.ParseInt(did, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	d.Data = append(d.Data, Device{
		Address: h,
		Dtype:   dtype,
		Did:     didInt,
		Status:  status,
		Rate:    1,
	})
	return nil
}

func (d *Devices) Add(address []byte, new Device) error {
	for _, device := range d.Data {
		if bytes.Equal(device.Address, address) {
			return ErrDuplicateDevice
		}
	}

	d.Data = append(d.Data, new)
	return nil
}

func (d *Devices) UpdateStatus(address []byte, user string, msg string) error {
	i := 0

	for _, device := range d.Data {
		if bytes.Equal(device.Address, address) {
			isTurnedOn := d.Data[i].Status

			if isTurnedOn {
				// Turning off the device
				if msg != "off" {
					log.Println("Already Turned On")
					break
				}
				if user != d.Data[i].User {
					log.Println("Turned Off by a Different User?!")
					break
				} else {
					d.Data[i].Status = false
					/* TODO: stop timer and store usage */
				}
			} else {
				// Turning on the device
				if msg != "on" {
					log.Println("Not Turned On Yet")
					break
				} else {
					d.Data[i].User = user
					d.Data[i].Status = true
					/* TODO: trigger timer */
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

func (d *Devices) Delete(address []byte) error {
	i := 0
	for _, device := range d.Data {
		if bytes.Equal(device.Address, address) {
			break
		}
		i++
	}

	if i < len(d.Data) {
		d.Data = append(d.Data[:i], d.Data[i+1:]...)
	} else {
		log.Println("No matching address")
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
