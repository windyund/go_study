package main

import "fmt"

// 值传递和指针传递
func changeValue(p int) {
	p = 10
}

// 指针传递，还存在多级指针情况
func changeValuePointer(p *int) {
	//指针
	*p = 10
}

func main() {
	var a int = 1
	changeValue(a)
	fmt.Println("a= ", a)

	changeValuePointer(&a)
	fmt.Print("a= ", a)

}
