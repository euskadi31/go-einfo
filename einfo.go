package einfo

import (
	"os"

	"github.com/euskadi31/go-einfo/internal"
)

func Info(format string, a ...interface{}) {
	internal.Info(os.Stdout, internal.EColorGood, format, a...)
}

func Warn(format string, a ...interface{}) {
	internal.Info(os.Stderr, internal.EColorWarn, format, a...)
}

func Error(format string, a ...interface{}) {
	internal.Info(os.Stderr, internal.EColorBad, format, a...)
}

func Begin(format string, a ...interface{}) {
	internal.Begin(os.Stdout, internal.EColorGood, format, a...)
}

func End(status bool, format string, a ...interface{}) {
	internal.DoEnd("eend", status, format, a...)
}

func Wend(status bool, format string, a ...interface{}) {
	internal.DoEnd("ewend", status, format, a...)
}
