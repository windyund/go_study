package main

import (
	"fmt"
	"reflect"
)

// 反射
func reflectNum(arg interface{}) {
	fmt.Println("type of arg is ", reflect.TypeOf(arg))
	fmt.Println("value of arg is ", reflect.ValueOf(arg))
}

func main() {
	var num int = 10
	reflectNum(num)
}
