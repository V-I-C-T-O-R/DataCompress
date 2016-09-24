package main

import (
	"flag"
	"github.com/V-I-C-T-O-R/DataCompress/decompress"
	"github.com/V-I-C-T-O-R/DataCompress/utils"
	"log"
	"path"
	"runtime"
)

var filePath string
var basePath string
var fileOutPath string

func init() {
	flag.StringVar(&filePath, "filePath", "output.json", "read to compress's file")
	flag.StringVar(&fileOutPath, "fileOutPath", "comeback.json", "read to output the compress's file")
	_, filename, _, _ := runtime.Caller(1)
	basePath = path.Join(path.Dir(filename), "example")
}
func main() {
	flag.Parse()
	filePath = basePath + "/" + filePath
	fileOutPath = basePath + "/" + fileOutPath
	data, err := decompress.DoDeCompress(filePath)
	if err != nil {
		log.Println("file decompress failed")
	} else {
		utils.WriteFile(data, fileOutPath)
		log.Println("file decompress complete")
	}
}
