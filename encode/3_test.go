package encode_test

import (
	"AniYa/encode"
	"bytes"
	"io/ioutil"
	"testing"
)

func Test3(t *testing.T) {
	key := []byte("testtesttest")
	testbyte, _ := ioutil.ReadFile("beacon.bin")
	beforebyte := testbyte
	encodedbyte := encode.Encode3(testbyte, key)
	afterbyte := encode.Decode3(encodedbyte, key)
	if bytes.Equal(afterbyte, beforebyte) {
		println("success")
	} else {
		t.Error("error")
	}
}
