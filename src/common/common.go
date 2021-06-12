package common

import "fmt"

// 包外引用的变量使用大驼峰，即以大写字母开头
// 常量类型为布尔、数字以及字符串类型
var PI float64 = 3.14

const PIEx = 3.1415926

// 包内引用的变量使用小驼峰，即以小写字母开头
var pi float64 = 3.14

// init()函数常量和变量初始化后自动执行
func init() {
	fmt.Print("init()\n")
}
