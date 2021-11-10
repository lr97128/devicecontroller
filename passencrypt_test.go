package main

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

const (
	FIRSTKEY  = "0123456789asdfgh"
	SECRET    = "0123456789asdfgh"
	CIPHPASS  = "yCjrCSibGdyU4WY0lQtcKw=="
	PLAINPASS = "liurui97128224"
)

func TestGetPlainPass(t *testing.T) {
	t.Log("BEGIN TO TEST DESCRYPT CIPHPASSWORD")
	ret, err := GetPlainPass(CIPHPASS, FIRSTKEY, SECRET)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(PLAINPASS, ret))
}

func TestGetCiphPass(t *testing.T) {
	t.Log("BEGIN TO TEST ENCRYPT PLAINPASSWORD")
	text := "liurui97128224"
	ret, err := GetCiphPass(text, FIRSTKEY, SECRET)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(CIPHPASS, ret))
}
