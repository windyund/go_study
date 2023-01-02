package main

import "fmt"

func main() {
	//1
	myMap := make(map[string]string)
	myMap["one"] = "1"
	myMap["two"] = "2"
	fmt.Println("map: ", myMap)
	//2
	var myMap2 map[int]string
	myMap2 = make(map[int]string)
	myMap2[1] = "test"

	//3
	var myMap3 = map[string]int{
		"1": 1,
		"2": 2,
	}
	fmt.Println("myMap3: ", myMap3)
}
