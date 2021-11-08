package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetDevices(path string) ([]Device, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	out, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var devices = []Device{}
	if err := json.Unmarshal(out, &devices); err != nil {
		return nil, err
	}
	return devices, nil
}
