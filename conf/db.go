// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var DeviceData Devices

var ErrDuplicateDevice = errors.New("Duplicate Device")

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
	h, err := Checksum(dtype, did)

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
		Usage:   make(map[string]int),
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

func (d *Devices) Update(address []byte, new Device) error {
	i := 0
	for _, device := range d.Data {
		if bytes.Equal(device.Address, address) {
			break
		}
		i++
	}

	if i < len(d.Data) {
		d.Data = append(d.Data[:i], d.Data[i+1:]...)
		d.Data = append(d.Data, new)
	} else {
		log.Println("No matching address")
	}

	return nil
}

func (d *Devices) Delete(address []byte) error {
	i := 0
	for i, device := range d.Data {
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
