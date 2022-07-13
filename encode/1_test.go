package encode_test

import (
	"AniYa/encode"
	"bytes"
	"io/ioutil"
	"testing"
)

func Test1(t *testing.T) {
	key := []byte("testtesttest")
	testbyte, _ := ioutil.ReadFile("beacon.bin")
	beforebyte := testbyte
	encodedbyte := encode.Encode1(testbyte, key)
	afterbyte := encode.Decode1(encodedbyte, key)
	if bytes.Equal(afterbyte, beforebyte) {
		println("success")
	} else {
		t.Error("error")
	}
}
