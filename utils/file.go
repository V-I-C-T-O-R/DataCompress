package utils

import (
	"io/ioutil"
	"os"
)

func ReadF(filePath string) []byte {
	file, err := os.Open(filePath)
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
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func WriteFile(data []byte, filename string) {
	var f *os.File
	var err1 error
	if CheckFileIsExist(filename) {
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666)
		check(err1)
	} else {
		f, err1 = os.Create(filename)
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
