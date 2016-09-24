package main

import (
	"flag"
	"github.com/V-I-C-T-O-R/DataCompress/compress"
	"path"
	"runtime"
)

var filePath string
var fileOutPath string
var basePath string

func init() {
	flag.StringVar(&fileOutPath, "fileOutPath", "output.json", "read to output the compress's file")
	flag.StringVar(&filePath, "filePath", "example.json", "read to compress's file")
	_, filename, _, _ := runtime.Caller(1)
	basePath = path.Join(path.Dir(filename), "example")
}
func main() {
	flag.Parse()
	filePath = basePath + "/" + filePath
	fileOutPath = basePath + "/" + fileOutPath
	compress.DoCompress(filePath, fileOutPath)
}
