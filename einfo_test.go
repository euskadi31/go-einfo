package einfo

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Output = tmpfile

	Info("test")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * test\n", string(actual))
}

func TestWarn(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Output = tmpfile

	Warn("test")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * test\n", string(actual))
}

func TestError(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Output = tmpfile

	Error("test")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * test\n", string(actual))
}

func TestBegin(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Output = tmpfile

	Begin("test")

	End(true, "test")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * test ...\n [ ok ]\n", string(actual))
}

func TestEnd(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Output = tmpfile

	Begin("fetch")

	End(false, "cannot fetch")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * fetch ...\n * cannot fetch\n                                                                 [ !! ]\n", string(actual))
}

func TestWend(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "einfo")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	Output = tmpfile

	Begin("fetch")

	Wend(false, "cannot fetch")

	actual, err := ioutil.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, " * fetch ...\n * cannot fetch\n                                                                 [ !! ]\n", string(actual))
}
