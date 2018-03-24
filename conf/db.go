// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var DeviceData Devices

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
	json.Unmarshal(raw, &d.data)
	return nil
}

func (d *Devices) Add(address []byte, new Device) error {
	d.data = append(d.data, new)

	for _, device := range d.data {
		log.Printf("%x\n", device.Address)
	}
	return nil
}

func (d *Devices) Update(address []byte, new Device) error {
	i := 0
	for _, device := range d.data {
		if bytes.Equal(device.Address, address) {
			break
		}
		i++
	}

	if i < len(d.data) {
		d.data = append(d.data[:i], d.data[i+1:]...)
		d.data = append(d.data, new)
	} else {
		log.Println("No matching address")
	}

	return nil
}

func (d *Devices) Delete(address []byte) error {
	i := 0
	for i, device := range d.data {
		if bytes.Equal(device.Address, address) {
			break
		}
		i++
	}

	if i < len(d.data) {
		d.data = append(d.data[:i], d.data[i+1:]...)
	} else {
		log.Println("No matching address")
	}

	return nil
}

func (d *Devices) Dump() error {
	bytes, err := json.Marshal(&d.data)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("bri-devices.json", bytes, 0666)
	return nil
}
