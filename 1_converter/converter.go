package converter

import "fmt"

func Converter() {
	const USDinEUR float64 = 0.85
	const USDinRUB float64 = 78.6

	var EUR float64 = 15

	EURtoRUB := EUR / USDinEUR * USDinRUB

	fmt.Println(EURtoRUB)
}
