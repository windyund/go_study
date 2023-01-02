package main

import "fmt"

func map_test() {

}

const (
	//iota 只能在const中使用
	BEIGING = iota
	HANGZHOU
	GUANGZHOU
)

func main() {
	var a = "1"
	fmt.Println("a = ", a)
	fmt.Printf("type of a %T\n", a)
	//该种声明，不支持全局变量，局部变量
	b := 1
	fmt.Println("b =", b)
	var c, d int = 1, 2
	fmt.Println("c=,d=", c, d)

	var (
		e    int    = 1
		st   string = "test"
		flag bool   = true
	)
	fmt.Println("e, st, flag ", e, st, flag)
	fmt.Println("beijing = ", BEIGING)
	fmt.Println("beijing = ", GUANGZHOU)
}
