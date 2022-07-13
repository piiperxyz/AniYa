package core

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func FileCopy(srcpath string, dstpath string) bool {
	srcFile, err := os.Open(srcpath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	buf := make([]byte, 1024)
	dstFile, err2 := os.Create(dstpath)
	if err2 != nil {
		fmt.Println(err)
		return false
	}
	for {
		// 从源文件读数据
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//写出去
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()
	return true
}

func RemoveSpecialCharactar(filedata string) string {
	filedata = strings.Replace(filedata, "\\x", "", -1)
	filedata = strings.Replace(filedata, "\"", "", -1)
	filedata = strings.Replace(filedata, " ", "", -1)
	filedata = strings.Replace(filedata, "\r\n", "", -1)
	filedata = strings.Replace(filedata, ";", "", -1)
	return filedata
}
