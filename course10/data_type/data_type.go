/*
 * @Descripttion:
 * @version:
 * @Author: Winston
 * @Date: 2020-04-16 09:35:06
 * @LastEditors: Winston
 * @LastEditTime: 2020-04-16 11:37:06
 */

// 基本数据类型

/*
整型
整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

类型	描述
uint8	无符号 8位整型 (0 到 255)
uint16	无符号 16位整型 (0 到 65535)
uint32	无符号 32位整型 (0 到 4294967295)
uint64	无符号 64位整型 (0 到 18446744073709551615)
int8	有符号 8位整型 (-128 到 127)
int16	有符号 16位整型 (-32768 到 32767)
int32	有符号 32位整型 (-2147483648 到 2147483647)
int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)


特殊整型
类型	描述
uint	32位操作系统上就是uint32，64位操作系统上就是uint64
int	32位操作系统上就是int32，64位操作系统上就是int64
uintptr	无符号整型，用于存放一个指针

注意：在使用int和 uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。

注意事项: 获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。
实际使用中，切片或 map 的元素数量等都可以用int来表示。在涉及到二进制传输、
读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint。
*/

//数字字面量法 以二进制、八进制或十六进制浮点数的格式定义数字

/*
浮点型
float32和float64

float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。
*/

/*
复数
complex64和complex128

复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
*/

/*
布尔值
以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值

注意：

1.布尔类型变量的默认值为false。
2.Go 语言中不允许将整型强制转换为布尔型.
3.布尔型无法参与数值运算，也无法与其他类型进行转换。
*/

/*
字符串

Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型
（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用UTF-8编码。
 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：
s1 := "hello"
s2 := "你好"

字符串转义符
\r	回车符（返回行首）
\n	换行符（直接跳到下一行的同列位置）
\t	制表符
\'	单引号
\"	双引号
\\	反斜杠

多行字符串
要定义一个多行字符串时，就必须使用反引号字符：
反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

字符串的常用操作

方法	                              介绍
len(str)	                          求长度
+或fmt.Sprintf	                      拼接字符串
strings.Split	                      分割
strings.contains	                  判断是否包含
strings.HasPrefix,strings.HasSuffix	  前缀/后缀判断
strings.Index(),strings.LastIndex()	  子串出现的位置
strings.Join(a[]string, sep string)	  join操作
*/

/*
byte和rune类型
组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
var a := '中'
var b := 'x'

1. uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
2. rune类型，代表一个 UTF-8字符。

当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，
也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
*/

package main

import (
	"fmt"
	"math"
)

// 遍历字符串 byte和rune类型演示
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	/*
		输出：
		104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)
		104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)

		因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，
		否则就会出现上面输出中第一行的结果。

		字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，
		所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	*/
}

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制
	fmt.Printf("%o \n", a) // 1010  占位符%b表示八进制
	fmt.Printf("%x \n", a) // 1010  占位符%b表示十六进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	//浮点演示
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//复数演示
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	//bool类型演示
	var t1 = true
	var f1 = false
	println(t1, f1)

	//字符串转义示例
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	//多行字符串
	s1 := `第一行
	第二行
	第三行
	`
	fmt.Println(s1)
}