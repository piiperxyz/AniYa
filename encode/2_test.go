package encode_test

import (
	"AniYa/encode"
	"bytes"
	"io/ioutil"
	"testing"
)

func Test2(t *testing.T) {
	key := []byte("testtesttest")
	testbyte, _ := ioutil.ReadFile("beacon.bin")
	beforebyte := testbyte
	encodedbyte := encode.Encode2(testbyte, key)
	afterbyte := encode.Decode2(encodedbyte, key)
	if bytes.Equal(afterbyte, beforebyte) {
		println("success")
	} else {
		t.Error("error")
	}
}
