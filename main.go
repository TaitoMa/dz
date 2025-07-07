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

func getNumber() int {
	var n int
	fmt.Scan(&n)
	return n
}

func calculateCurrency(x int, currency1 string, currency2 string) {

}
