package terminal

import (
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

var (
	isColor = false
	term    = ""
	terms   = [36]string{
		"Eterm",
		"ansi",
		"con132x25",
		"con132x30",
		"con132x43",
		"con132x60",
		"con80x25",
		"con80x28",
		"con80x30",
		"con80x43",
		"con80x50",
		"con80x60",
		"cons25",
		"console",
		"cygwin",
		"dtterm",
		"gnome",
		"konsole",
		"kterm",
		"linux",
		"linux-c",
		"mlterm",
		"putty",
		"rxvt",
		"rxvt-cygwin",
		"rxvt-cygwin-native",
		"rxvt-unicode",
		"screen",
		"screen-bce",
		"screen-w",
		"screen.linux",
		"vt100",
		"vt220",
		"wsvt25",
		"xterm",
		"xterm-debian",
	}
)

func IsColor(f *os.File) bool {
	if !isatty.IsTerminal(f.Fd()) {
		return false
	}

	if isColor {
		return true
	}

	if term == "" {
		term = os.Getenv("TERM")

		if term == "" {
			return false
		}
	}

	if strings.Contains(term, "color") {
		isColor = true

		return isColor
	}

	for _, termName := range terms {
		if strings.Compare(termName, term) == 0 {
			isColor = true

			return isColor
		}
	}

	return false
}
