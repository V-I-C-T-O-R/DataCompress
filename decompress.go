package main

import (
	"flag"
	"github.com/V-I-C-T-O-R/DataCompress/decompress"
	"path"
	"runtime"
)

var filePath string
var basePath string

func init() {
	flag.StringVar(&filePath, "filePath", "output.json", "read to compress's file")
	_, filename, _, _ := runtime.Caller(1)
	basePath = path.Join(path.Dir(filename), "example")
}
func main() {
	flag.Parse()
	filePath = basePath + "/" + filePath
	decompress.DoDeCompress(filePath)
}
