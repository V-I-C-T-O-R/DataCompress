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
	if !CheckFileIsExist(filename) {
		_, err := os.Create(filename)
		check(err)
	}
	err := ioutil.WriteFile(filename, data, 0644)
	check(err)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
