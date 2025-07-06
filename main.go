package main

import (
	"fmt"
)

func main() {
	const USDinEUR float64 = 0.85
	const USDinRUB float64 = 78.6

	const EUR float64 = 15

	const EURtoRUB = EUR / USDinEUR * USDinRUB

	fmt.Println(EURtoRUB)
}
