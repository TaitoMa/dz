package main

import (
	"fmt"
)

func main() {
	const USDinEUR float64 = 0.85
	const USDinRUB float64 = 78.6

	const EURtoRUB = USDinRUB / USDinEUR

	fmt.Println(EURtoRUB)
}
