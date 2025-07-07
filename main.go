package main

import (
	"fmt"
)

func main() {
	const USDinEUR float64 = 0.85
	const USDinRUB float64 = 78.6

	const EURtoRUB = USDinRUB / USDinEUR

	fmt.Println(EURtoRUB)

	number := getNumber()

	fmt.Println(number)
}

func getNumber() float64 {
	var n float64
	fmt.Scan(&n)
	return n
}

func calculateCurrency(x float64, currency1 string, currency2 string) {

}
