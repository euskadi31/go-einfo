package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

// WindowSize struct.
type WindowSize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// GetWindowSize from term.
func GetWindowSize(f *os.File) (*WindowSize, error) {
	ws := &WindowSize{}

	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		f.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	// nolint: gosec
	if int(retCode) == -1 {
		return nil, errno
	}

	return ws, nil
}
