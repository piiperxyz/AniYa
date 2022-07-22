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

func FileCopy(srcPath string, dstPath string) bool {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	buf := make([]byte, 1024)
	dstFile, err2 := os.Create(dstPath)
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

func RemoveSpecialCharacter(fileData string) string {
	fileData = strings.Replace(fileData, "\\x", "", -1)
	fileData = strings.Replace(fileData, "\"", "", -1)
	fileData = strings.Replace(fileData, " ", "", -1)
	fileData = strings.Replace(fileData, "\r\n", "", -1)
	fileData = strings.Replace(fileData, ";", "", -1)
	return fileData
}
