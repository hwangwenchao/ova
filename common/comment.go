package common

import "fmt"

func init() {
	fmt.Print("comment init()\n")
}

func Comment() {
	fmt.Print("Commnet()\n")
}

// 包内引用的函数使用小驼峰，即以小写字母开头
func common() {
	fmt.Print("common()\n")
}

// 包外引用的函数使用大驼峰，即以大写字母开头。|否则，在运行时会出错
func Log() {
	fmt.Print("log()\n")
}
