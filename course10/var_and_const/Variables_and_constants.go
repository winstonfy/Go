/* 
标识符
break        default      func         interface    select
//case         defer        go           map          struct
//chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var

37个保留字
Constants:    true  false  iota  nil

Types:    int  int8  int16  int32  int64  
			uint  uint8  uint16  uint32  uint64  uintptr
			float32  float64  complex128  complex64
			bool  byte  rune  string  error

Functions:   make  len  cap  new  append  copy  close  delete
			complex  real  imag
			panic  recover
*/
package main


import (
	"fmt"
)
//变量 
//变量声明
// var 变量名 变量类型
var name1 string
var age1 int
var isOk bool

//批量声明变量
var (
    a string
    b int
    c bool
    d float32
)

//变量的初始化
//var 变量名 类型 = 表达式
var name2 string = "Winston"
var age2 int = 18
//一次性初始化多个值
var name3, age3 = "Winston", 20

// 类型推导
//将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化
var name4 = "Winston"
var age4 = 18

//匿名变量
// 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示
func foo() (int, string) {
	return 10, "Winston"
}


// 常量
//把var换成了const，常量在定义的时候必须赋值。
const pi1 = 3.1415
const e1 = 2.7182
// 常量的批量声明
const (
    pi2 = 3.1415
    e2 = 2.7182
)

//const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
    n1 = 100
    n2
    n3
)

// iota 面试常考
/*
iota是go语言的常量计数器，只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次
(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
*/
const (
	m1 = iota //0
	m2        //1
	m3        //2
	m4        //3
)
//几个常见的iota示例:
const (
	h1 = iota //0
	h2        //1
	_
	h4        //3
)

const (
	i1 = iota //0
	i2 = 100  //100
	i3 = iota //2
	i4        //3
)
const i5 = iota //0

const (
	_  = iota 
	/*KB 1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024*/
	KB = 1 << (10 * iota)
	/*MB xxx*/
	MB = 1 << (10 * iota)
	/*GB xxx*/
	GB = 1 << (10 * iota)
	/*TB xxx*/
	TB = 1 << (10 * iota)
	/*PB xxx*/
	PB = 1 << (10 * iota)
)
//多个iota定义在一行
const (
	a1, b1 = iota + 1, iota + 2 //1,2
	c1, d1                      //2,3
	j1, k1                      //3,4
)

func main() {
	//短变量声明
	//在函数内部，可以使用更简略的 := 方式声明并初始化变量
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(n,m)

	//匿名变量使用
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	// 函数外的每个语句都必须以关键字开始（var、const、func等）
	// :=不能使用在函数外。
	// _多用于占位，表示忽略值。
}