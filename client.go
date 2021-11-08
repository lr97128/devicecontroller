package main

import (
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func GetClientWithPassword(dev Device) (c *ssh.Client, err error) {
	config := &ssh.ClientConfig{
		User:            dev.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(dev.PlainPass)},
		Timeout:         time.Second * 10,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return ssh.Dial("tcp", dev.Host, config)
}

func GetClientWithKey(dev Device) (c *ssh.Client, err error) {
	if dev.Key == "" {
		dev.Key = os.ExpandEnv("$HOME/.ssh/id_rsa")
	}
	file, err := os.Open(dev.Key)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	res, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	pk, err := ssh.ParsePrivateKey(res)
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(pk)},
		Timeout:         time.Second * 10,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return ssh.Dial("tcp", dev.Host, config)
}

func Execute(c *ssh.Client, cmds []string) (results []string, error) {
	sess, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()
	for _, cmd := range cmds {
		results := append(results, string(sess.CombinedOutput(cmd)))
	}
	// return sess.CombinedOutput(cmd)
}
