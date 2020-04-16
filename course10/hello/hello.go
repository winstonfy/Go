// 第一个go程序

package main // 声明 main 包，表明当前是一个可执行程序

import ( 
	"fmt"  // 导入内置 fmt 包
)

func main() { // main函数，是程序执行的入口
	fmt.Println("Let's Go")  // 在终端打印 Hello World!
}

// go build表示将源代码编译成可执行文件。

//还可以使用-o参数来指定编译后得到的可执行文件的名字。