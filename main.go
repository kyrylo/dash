// Print a colourful full width dash.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

const (
	dashCh     = "-"
	greenStart = "\033[32m"
	greenEnd   = "\033[0m"
)

type window struct {
	row uint16
	col uint16
}

func getWidth() uint {
	win := &window{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(win)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(win.col)
}

func main() {
	var str strings.Builder
	w := int(getWidth())

	if len(os.Args) > 1 {
		lines, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		str.WriteString(greenStart)

		for i := 0; i < lines; i++ {
			str.WriteString(strings.Repeat(dashCh, w))
		}

		str.WriteString(greenEnd)
	} else {
		str.WriteString(greenStart)
		str.WriteString(strings.Repeat(dashCh, w))
		str.WriteString(greenEnd)
	}

	fmt.Printf(str.String())
}
