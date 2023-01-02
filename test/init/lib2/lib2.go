package lib2

import "fmt"

// 方法名大写表示 public,对外可见，否则只有本包可见
func Lib2Test() {
	fmt.Print("Lib2Test running!")
}

func init() {
	fmt.Println("lib2  init ....")
}
