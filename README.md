`基于json数据的数组数据存储的压缩工具`
------------------------------------
DataCompress针对大数据中json数组数据存储冗余的情况，进行数据压缩存储，并进行解析。实现json数据的压缩和解析功能

###使用说明：
####1.压缩步骤
    默认文件提取位置位于项目的example目录下，即可以将需要压缩的json文件存放在这里，然后终端切换到项目根目录下，运行
    compress.go文件，如：
    go run compress.go -filePath="example.json" -fileOutPath="output.json"
    其中，`filePath`表示的需要压缩的json文件名称，`fileOutPath`表示压缩后产生的新文件的名称
压缩前后的示例如图：

![](https://github.com/V-I-C-T-O-R/DataCompress/blob/master/image/example.png) ![](https://github.com/V-I-C-T-O-R/DataCompress/blob/master/image/output.png) 
####2.还原步骤
    默认文件提取位置位于项目的example目录下，即可以将需要还原的json文件存放在这里，然后终端切换到项目根目录下，运行
    decompress.go文件，如：
    go run decompress.go -filePath="output.json" -fileOutPath="comeback.json"
    其中，`filePath`表示的需要还原的json文件名称

=====
####`须知:`
  `待还原的json文件必须是使用项目压缩之后的json文件`
###`补充:`
除了上述的使用方式外，还可以直接调用compress文件夹中compress.go文件的接口直接进行压缩，调用decompress文件夹中的decompress文件夹中的decompress.go文件的接口进行还原输出.
    例如：compress/compress.go中的DoCompress(file string, outputstring)函数直接传入待解析和输出的文件的绝对地址，即可进行压缩。
    DoStreamCompress(b []byte, output string)函数直接传入[]byte数据进行解析，output绝对路径进行存储
    decompress/decompress.go中的func DoDeCompress(file string) (data []byte, err error)函数直接传入待解析文件的绝对地址，即可进行还原出[]byte数据。DoDeCompressFromData(b []byte) (data []byte, err error)函数直接传入[]byte数据进还原析，即可进行还原出[]byte数据。(注：仅支持项目压缩产生的json数据形式)
