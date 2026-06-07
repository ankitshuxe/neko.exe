//go:build windows

package cmd

import (
	"os"

	"golang.org/x/sys/windows"
)

func setupTerminal() {
	stdout := windows.Handle(os.Stdout.Fd())
	var mode uint32
	windows.GetConsoleMode(stdout, &mode)
	windows.SetConsoleMode(stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	// 65001 is CP_UTF8
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	setConsoleOutputCP := kernel32.NewProc("SetConsoleOutputCP")
	setConsoleOutputCP.Call(uintptr(65001))
}
