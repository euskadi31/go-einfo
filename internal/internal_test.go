package internal

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	color := Color(os.Stdout, EColorGood)

	assert.Equal(t, "", color)
}

func TestInfo(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Info(tmpfile, EColorGood, "test")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * test\n", string(actual))
}

func TestBegin(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Begin(tmpfile, EColorGood, "test")

	DoEnd(tmpfile, "eend", true, "test")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * test ...\n [ ok ]\n", string(actual))
}
