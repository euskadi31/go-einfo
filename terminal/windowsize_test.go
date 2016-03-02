package terminal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWindowSize(t *testing.T) {
	info, err := GetWindowSize(os.Stdout)

	assert.Nil(t, info)
	assert.Error(t, err, "operation not supported by device")
}
