package main

import (
	"fmt"
	"go_study/test/init/lib1"
	//匿名,可以不使用该包下的方法
	_ "go_study/test/init/lib1"
	//将方法导入到当前层级包中，注意方法重名情况
	// my 别名
	. "go_study/test/init/lib2"
)

// 类似构造方法
func init() {
	fmt.Println("main init ...")
}

func main() {
	lib1.Lib1Test()
	Lib2Test()
}
