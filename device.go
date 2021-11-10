package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Device struct {
	Host      string `json:"host"`
	Username  string `json:"username"`
	CiphPass  string `json:"ciphpass"`
	PlainPass string
	Secret    string   `json:"secret"`
	Key       string   `json:"key"`
	UseKey    bool     `json:"usekey"`
	Commands  []string `json:"commands"`
	Client    *ssh.Client
}

func (d *Device) InitDevice(firstKey string) error {
	plainPass, err := GetPlainPass(d.CiphPass, firstKey, d.Secret)
	if err != nil {
		return err
	}
	d.PlainPass = plainPass
	client, err := GetClientWithPassword(*d)
	if err != nil {
		return err
	}
	d.Client = client
	return nil
}

func (d *Device) Execute() ([]string, error) {
	cmds := d.Commands
	var result []string
	for _, cmd := range cmds {
		ret, err := Execute(d.Client, cmd)
		if err != nil {
			return nil, err
		}
		result = append(result, ret)
	}
	return result, nil
}

func (d *Device) Close() error {
	if d.Client != nil {
		if err := d.Client.Close(); err != nil {
			return err
		}
	}
	return nil
}
