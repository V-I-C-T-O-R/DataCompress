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
	if checkFileIsExist(filename) {
		f, err1 = os.OpenFile(file, os.O_APPEND, 0666)
		check(err1)
	} else {
		f, err1 = os.Create(file)
		check(err1)
	}
	defer f.Close()
	_, err1 = f.Write(data)
	check(err1)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
