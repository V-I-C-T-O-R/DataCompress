package main

import (
	"flag"
	"github.com/V-I-C-T-O-R/DataCompress/compress"
)

var FilePath string
var FileOutPath string

func init() {
	flag.StringVar(&FileOutPath, "FileOutPath", "output.json", "read to output the compress's file")
	flag.StringVar(&FilePath, "FilePath", "example.json", "read to compress's file")
}
func main() {
	flag.Parse()
	compress.DoCompress(FilePath, FileOutPath)
}
