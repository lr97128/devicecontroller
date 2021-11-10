package main

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetClientWithPassword(t *testing.T) {
	FIRSTKEY := "0123456789asdfgh"
	t.Log("test to get client and execute command")
	dev := Device{
		Host:     "10.199.99.140:22",
		Username: "root",
		CiphPass: "yCjrCSibGdyU4WY0lQtcKw==",
		UseKey:   false,
		Commands: []string{"uname -r"},
	}
	t.Log("Before Init, device:", dev)
	dev.InitDevice(FIRSTKEY)
	t.Log("After Init, device:", dev)
	c, err := GetClientWithPassword(dev)
	assert.NilError(t, err)
	defer c.Close()
	s, err := c.NewSession()
	assert.NilError(t, err)
	out, err := s.CombinedOutput(dev.Commands[0])
	assert.NilError(t, err)
	ret := string(out)
	want := "3.10.0-1160.el7.x86_64"
	assert.Check(t, is.Equal(want, ret))
}
