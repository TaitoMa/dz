package main

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()

	bmw := BMW{}
	zhiga := Zhiga{}
	sayHello(zhiga)
	sayHello(zhiga)
	sayHello(zhiga)
	sayHello(bmw)

	fmt.Println(time.Since(before))
}
