package encode

import (
	"fmt"
	sgn "github.com/EgeBalci/sgn/pkg"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func Sgn(path string) {
	// First open some file
	file, err := ioutil.ReadFile(path)
	if err != nil { // check error
		fmt.Println(err)
	}
	// Create a new SGN encoder
	encoder := sgn.NewEncoder()
	//sgn循环次数为5-10次
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(5)
	randomNum += 5
	encoder.EncodingCount = randomNum
	// Set the proper architecture
	encoder.SetArchitecture(64)
	// Encode the binary
	encodedBinary, err := encoder.Encode(file)
	if err != nil {
		fmt.Println(err)
	}
	// Print out the hex dump of the encoded binary
	ioutil.WriteFile("shellcode", encodedBinary, os.ModePerm)
	println("sgn done!")
}
