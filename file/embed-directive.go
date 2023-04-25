package main

import (
	"embed"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

// //go:embed 这么设计的动机是什么？

// 见 https://www.cnblogs.com/rxbook/p/15753862.html

// embed的基本语法
// 基本语法非常简单，首先导入embed包，然后使用指令//go:embed 文件名 将对应的文件或目录结构导入到对应的变量上

// 注意事项
// 1. 在使用//go:embed指令的文件都需要导入 embed包。 例如，以下例子 没有导入embed包，即不会正常运行
// 2. //go:embed指令只能用在包一级的变量中，不能用在函数或方法级别
// 3.1 当包含目录时，它不会包含以“.”或““开头的文件。但是如果使用通配符，比如dir/*，它将包含所有匹配的文件，即使它们以“."或""开头。
// 3.2 请记住，在您希望在Web服务器中嵌入文件但不允许用户查看所有文件的列表的情况下，包含Mac OS的.DS_Store文件可能是一个安全问题。
// 3.3 出于安全原因，Go在嵌入时也不会包含符号链接或上一层目录。
