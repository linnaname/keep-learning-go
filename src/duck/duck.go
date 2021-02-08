package main

import "fmt"

type Biology interface {
	sayhi()
}

type Man struct {
	name string
	age  int
}

type Monster struct {
	name string
	age  int
}

func (this Man) sayhi() { // 实现抽象方法1
	fmt.Printf("Man[%s, %d] sayhi\n", this.name, this.age)
}

func (this Monster) sayhi() { // 实现抽象方法1
	fmt.Printf("Monster[%s, %d] sayhi\n", this.name, this.age)
}

func WhoSayHi(i Biology) {
	i.sayhi()
}

func main() {
	b := test()
	m, _ := b.(Monster)
	fmt.Println(m)
}

func test() Biology {
	var b Biology
	b = Man{}
	return b
}
