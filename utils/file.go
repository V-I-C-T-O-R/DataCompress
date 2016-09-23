package utils

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

var basePath string

func init() {
	_, filename, _, _ := runtime.Caller(1)
	basePath = path.Join(path.Dir(filename), "..", "example")
}
func ReadF(filePath string) []byte {
	file, err := os.Open(basePath + "/" + filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return content
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func WriteFile(data []byte, filename string) {
	var f *os.File
	var err1 error
	file := basePath + "/" + filename
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(file, os.O_APPEND, 0666) //打开文件
		check(err1)
	} else {
		f, err1 = os.Create(file) //创建文件
		check(err1)
	}
	defer f.Close()
	_, err1 = f.Write(data) //写入文件(字节数组)
	check(err1)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
