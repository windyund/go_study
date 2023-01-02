package main

import "fmt"

//1.defer 方法测试，类似finally执行，支持多个defer, 按照栈的执行顺序
//2.defer 与 return 谁先执行, return 先执行

func defer_fun() string {
	fmt.Println("defer fun")
	return "hhh"
}
func defer_fun2() string {
	defer fmt.Println("defer fun2")
	return defer_fun()
}

func main() {
	//defer  defer_fun()
	defer defer_fun2()
}
