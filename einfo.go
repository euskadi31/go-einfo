package einfo

import (
	"os"

	"github.com/euskadi31/go-einfo/internal"
)

// Info print log to standard output. Arguments are handled in the manner of fmt.Print.
func Info(format string, a ...interface{}) {
	internal.Info(os.Stdout, internal.EColorGood, format, a...)
}

// Warn print log to standard output. Arguments are handled in the manner of fmt.Print.
func Warn(format string, a ...interface{}) {
	internal.Info(os.Stderr, internal.EColorWarn, format, a...)
}

// Error print log to standard output. Arguments are handled in the manner of fmt.Print.
func Error(format string, a ...interface{}) {
	internal.Info(os.Stderr, internal.EColorBad, format, a...)
}

// Begin a new message to print log to standard output. Arguments are handled in the manner of fmt.Print.
func Begin(format string, a ...interface{}) {
	internal.Begin(os.Stdout, internal.EColorGood, format, a...)
}

// End the message started with Begin to print log to standard output. Arguments are handled in the manner of fmt.Print.
func End(status bool, format string, a ...interface{}) {
	internal.DoEnd(os.Stdout, "eend", status, format, a...)
}

// Wend end the message with warnning level started with Begin to print log to standard output. Arguments are handled in the manner of fmt.Print.
func Wend(status bool, format string, a ...interface{}) {
	internal.DoEnd(os.Stdout, "ewend", status, format, a...)
}
