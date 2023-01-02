package main

import "fmt"

type myInt int

type Book struct {
	name string
	age  int
}

type ChildrenBook struct {
	Book
	publisher string
}

func (this *Book) getName() string {
	return this.name
}
func (this *Book) setName(name string) {
	this.name = name
}
func (this *Book) show() {
	fmt.Println("object is ", this)
}

func main() {
	var book Book
	book.name = "三国演义"
	book.age = 18
	fmt.Println(book)
	book1 := Book{
		name: "测试",
		age:  12,
	}
	book1.show()
	book1.setName("张无忌")
	book1.show()
	//fmt.Println("book1 is ", book1)
	var age myInt = 10
	fmt.Printf("type of age is %T", age)

	child := ChildrenBook{Book{"金庸", 12}, "中信出版"}
	fmt.Println("child is ", child)
}
