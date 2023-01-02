package main

import "fmt"

func return_one(a int, b int) int {
	return a + b
}

// 返回值匿名
func return_two(a int, b int) (int, int) {
	return a, b
}

// 返回值非匿名
func return_two_name(a int, b int) (r1 int, r2 int) {
	r1 = a
	r2 = b
	return r1, r2
}

func main() {
	c := return_one(1, 2)
	fmt.Print("c= ", c)
}
