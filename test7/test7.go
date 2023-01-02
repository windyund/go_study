package main

import "fmt"

type AnimalIF interface {
	Sleep()
	Walk()
	getName() string
}

type Dog struct {
	name string
}

func (this *Dog) Walk() {
	fmt.Print("dog walk....")
}

func (this *Dog) Sleep() {
	fmt.Print("dog sleep....")
}
func (this *Dog) getName() string {
	return this.name
}

type Cat struct {
	name string
}

func (this *Cat) Walk() {
	fmt.Print("cat walk....")
}
func (this *Cat) Sleep() {
	fmt.Print("cat sleep....")
}
func (this *Cat) getName() string {
	return this.name
}

// 多态测试
func main() {
	var animal AnimalIF
	animal = &Dog{"dog"}
	animal.Walk()

	animal = &Cat{"cat"}
	animal.Walk()
}
