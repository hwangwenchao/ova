/*
package main

import (
	"flag"
	. "fmt"
)

var mode = flag.String("mode", "", "proccess mode")

// main包中常量和变量初始化完成后再执行init()函数，然后再执行main()函数
// init()函数的作用就是对当前包中的变量进行初始化、检查/修复程序状态、注册以及运行一次计算
// init()函数由编译器自动调用，不可主动调用|每个包可有多个init()函数
func init() {
	fmt.Print("main init()")
}

func testFunc() {
	fmt.Println("testFunc()\n")
}

func main() {
	fmt.Print("hello world\n")

	flag.Parse()

	fmt.Println("mode:", *mode)

	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素

	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	testFunc()

	Println(("-------------end----------------"))
}
*/

package main

// import . 可以导入指定包，后续使用时可不用再加包名的访问方式。但是会造成相同函数、常量名称冲突
import (
	"fmt"
	"ova/common"
	"ova/msssql"
	"reflect"
	"unsafe"
)

type myint int32

// 全局变量必须以var开头，不可简写
var PI float32 = 3.1415926

func main() {
	fmt.Print("-------------start---------------\n")
	fmt.Printf("PI=%f\n", common.PI)
	var d bool = true
	fmt.Print("d:", d)

	var dd uint64 = 1
	var ddd int = 3        // int后不跟位数，则会根据操作系统来判断int的长度。
	var dddd float32 = 1.2 // float只有32或者64位
	var ddddd rune = 1     // 相当于int32
	fmt.Printf("size:%d\n", unsafe.Sizeof(dd))
	fmt.Printf("size:%d\n", unsafe.Sizeof(ddd))
	fmt.Printf("size:%d\n", unsafe.Sizeof(dddd))
	fmt.Printf("size:%d\n", unsafe.Sizeof(ddddd))

	var hasStop bool      // 默认值为false
	var intStop int32     // 默认为0
	var strStop string    // 默认为空
	var comStop complex64 // 默认值为0

	fmt.Print("hasStop:", hasStop)
	fmt.Print("\n")
	fmt.Print("intStop:", intStop)
	fmt.Print("\n")
	fmt.Print("strStop:", strStop)
	fmt.Print("\n")
	fmt.Print("comStop:", comStop)
	fmt.Print("\n")

	var alisVar myint
	fmt.Print("alis:", alisVar)
	fmt.Print("\n")
	fmt.Print("type:", reflect.TypeOf(alisVar))
	fmt.Print("\n")
	fmt.Print("size:", unsafe.Sizeof(alisVar))
	fmt.Print("\n")

	// 类型不同不可进行算数等操作
	// 提示报错：类型不匹配

	//fmt.Print("intStop+alisVar", intStop+alisVar)

	s := "DEMACIA"
	fmt.Print("s:", s)
	fmt.Print("\n")

	fmt.Print("type transform:", uint64(ddd))
	fmt.Print("\n")

	fmt.Print("PI:", common.PI)
	fmt.Print("\n")
	fmt.Print("PIEx:", common.PIEx)
	fmt.Print("\n")
	fmt.Print("type of PIEx:", reflect.TypeOf(common.PIEx))
	fmt.Print("\n")

	fmt.Print("connect database...\n")
	db := msssql.Msssql{
		DataSource: "hddf021.my3w.com",
		Database:   "hds1sf21_db",
		IsWindows:  false,
		Sa: msssql.SA{
			User:     "hdssdf021",
			Password: "",
		},
	}

	// connect database
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
	}
	defer db.Close()
	fmt.Print("close database")
	fmt.Print("\n")

	fmt.Print("--------------end----------------\n")
}
