package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/euskadi31/einfo-go/colors"
	"github.com/euskadi31/einfo-go/terminal"
)

type EColor int

const (
	EColorNormal EColor = iota
	EColorGood
	EColorWarn
	EColorBad
	EColorHilite
	EColorBracket
)

const (
	EOk    = "ok"
	ENotOk = "!!"
)

var (
	colorsMap = map[EColor]string{
		EColorNormal:  colors.ResetAll,
		EColorGood:    colors.Green,
		EColorWarn:    colors.Yellow,
		EColorBad:     colors.Red,
		EColorHilite:  colors.Cyan,
		EColorBracket: colors.Blue,
	}
	term         = ""
	termIsCons25 = false
)

func Color(f *os.File, color EColor) string {
	if !terminal.IsColor(f) {
		return ""
	}

	return colorsMap[color]
}

func info(f *os.File, color EColor, format string, a ...interface{}) int {
	length := 0

	n, _ := fmt.Fprintf(f, " %s*%s ", Color(f, color), Color(f, EColorNormal))

	length += n

	//@TODO indent

	n, _ = fmt.Fprintf(f, format, a...)

	length += n

	return length
}

func end(f *os.File, col int, color EColor, msg string) {
	cols := 0

	if msg == "" {
		return
	}

	windowSize, err := terminal.GetWindowSize(f)
	if err != nil {
		cols = 80
	}

	cols = int(windowSize.Col) - (len(msg) + 5)

	if term == "" {
		term = os.Getenv("TERM")

		if term != "" && strings.Compare(term, "cons25") == 0 {
			termIsCons25 = true
		} else {
			termIsCons25 = false
		}
	}

	if termIsCons25 {
		cols--
	}

	if cols > 0 && terminal.IsColor(f) {
		fmt.Fprintf(
			f,
			"%s%s %s[%s %s %s]%s\n",
			"\x1b[A",
			fmt.Sprintf("\x1b[%dC", cols),
			Color(f, EColorBracket),
			Color(f, color),
			msg,
			Color(f, EColorBracket),
			Color(f, EColorNormal),
		)
	} else {
		if col > 0 {
			for i := 0; i < cols-col; i++ {
				fmt.Fprint(f, " ")
			}
		}

		fmt.Fprintf(f, " [ %s ]\n", msg)
	}
}

func Info(f *os.File, color EColor, format string, a ...interface{}) {
	info(f, color, format, a...)

	fmt.Fprint(f, "\n")
}

func Begin(f *os.File, color EColor, format string, a ...interface{}) {
	info(f, color, format, a...)

	fmt.Fprint(f, " ...")
	fmt.Fprint(f, "\n")
}

func DoEnd(cmd string, status bool, format string, a ...interface{}) {
	col := 0
	f := os.Stdout

	if format != "" && status == false {
		f = os.Stderr

		if strings.Compare(cmd, "ewend") == 0 {
			col = info(f, EColorWarn, format, a...)
		} else {
			col = info(f, EColorBad, format, a...)
		}

		n, _ := fmt.Fprint(f, "\n")

		col += n
	}

	color := EColorGood
	msg := EOk

	if status == false {
		color = EColorBad
		msg = ENotOk
	}

	end(f, col, color, msg)
}
