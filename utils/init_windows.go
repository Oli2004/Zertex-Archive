// +build windows

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

func init() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

func AntiCrack() {
	programs := []string{"HTTPDebuggerSvc.exe", "Fiddler.exe", "x64dbg.exe", "x32dbg.exe", "PETools.exe", "MegaDumper.exe", "ExtremeDumper-x86.exe", "ExtremeDumper.exe"}
	cmd := exec.Command("tasklist.exe", "/fo", "csv", "/nh")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, _ := cmd.Output()
	for _, program := range programs {
		if strings.Contains(string(out), program) {
			fmt.Println("Nice skid tools..\nClosing Program!")
			time.Sleep(2 * time.Second)
			os.Exit(3)
		}
	}
}

func SetConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	return int(r), err
}
