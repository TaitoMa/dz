package main

import (
	"fmt"
)

type Auto interface {
	Hello()
}
type BMW struct{}

type Zhiga struct{}

func (b BMW) Hello() {
	fmt.Println("Hello, i'm BMW")
}

func (b Zhiga) Hello() {
	fmt.Println("Hello, i'm Zhiga")
}

func sayHello(auto Auto) {
	auto.Hello()
}

//func main() {
//	bmw := BMW{}
//	zhiga := Zhiga{}
//	sayHello(zhiga)
//	sayHello(bmw)
//}
