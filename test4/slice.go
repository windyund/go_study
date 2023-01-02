package main

import "fmt"

func main() {
	myArray := [10]int{1, 2, 3}
	for index, value := range myArray {
		fmt.Println("index =", index, "value = ", value)
	}
	//动态数组，是slice 切片，传参是指针传递
	dynamic := []int{1, 2, 3}
	printArray(dynamic)
	printArray(dynamic)

	//第二种声明方式
	var slice1 []int
	//分配空间大小，指定数据类型
	slice1 = make([]int, 3, 5)
	slice1[0] = 1

	fmt.Println("slice1 len= ", len(slice1), "cap= ", cap(slice1))
	if slice1 == nil {
		fmt.Println("slice1 is nil")
	} else {
		fmt.Println("slice1 是有空间的")
	}
	slice2 := slice1[0:1]
	printArray(slice2)
}
func printArray(arr []int) {
	for _, value := range arr {
		fmt.Println("value= ", value)
	}
	arr[0] = 10
}
