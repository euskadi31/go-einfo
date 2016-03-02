package terminal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsColor(t *testing.T) {
	assert.False(t, IsColor(os.Stdout))
}
