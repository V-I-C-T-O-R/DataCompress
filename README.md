`基于json数据的数组数据存储的压缩工具`
------------------------------------
DataCompress针对大数据中json数组数据存储冗余的情况，进行数据压缩存储，并进行解析。实现json数据的压缩和解析功能

###使用说明：
####1.压缩步骤
                默认文件提取位置位于项目example目录下，即可以将需要压缩的json文件存放在这里，然后终端切换到项目根目录下，运行
                compress.go文件，如：
                go run compress.go -filePath="example.json" -fileOutPath="output.json"
                其中，`filePath`表示的需要压缩的json文件名称，`fileOutPath`表示压缩后产生的新文件的名称
                压缩前后的示例如图：

![](https://github.com/V-I-C-T-O-R/DataCompress/blob/master/image/example.png) ![](https://github.com/V-I-C-T-O-R/DataCompress/blob/master/image/output.png) 
####2.还原步骤
