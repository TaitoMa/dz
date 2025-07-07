package main

import (
	"fmt"
)

const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
)

func main() {

	currencyFrom := getCurrency("")
	money := getNumber()
	currencyTo := getCurrency(currencyFrom)

	result := calculateCurrency(money, currencyFrom, currencyTo)

	fmt.Println(result)
}

func getCurrency(choosenCurrency string) string {
	fmt.Print("Введите валюту (USD/EUR/RUB): ")
	for {
		var currency string
		_, err := fmt.Scan(&currency)
		if err != nil {
			continue
		}
		if currency == choosenCurrency {
			fmt.Println("Нельзя перевести из одной валюты в такую-же")
			continue
		}
		if currency != "USD" && currency != "EUR" && currency != "RUB" {
			fmt.Println("Такой валюты нет")
			continue
		}

		return currency
	}
}

func getNumber() float64 {
	for {
		fmt.Println("Введите кол-во валюты: ")
		var n float64
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Нужно ввести число")
			continue
		}
		return n
	}
}

func calculateCurrency(x float64, currency1 string, currency2 string) float64 {
	const USDinEUR float64 = 0.85
	const USDinRUB float64 = 78.6

	const EURtoRUB = USDinRUB / USDinEUR

	var result float64

	switch {
	case currency1 == USD && currency2 == EUR:
		result = x * USDinEUR
	case currency1 == USD && currency2 == RUB:
		result = x * USDinRUB
	case currency1 == EUR && currency2 == USD:
		result = x / USDinEUR
	case currency1 == EUR && currency2 == RUB:
		result = x * EURtoRUB
	case currency1 == RUB && currency2 == EUR:
		result = x / EURtoRUB
	case currency1 == RUB && currency2 == USD:
		result = x / USDinRUB
	}
	return result
}

/* NOTES */
// errors.New()
// Если нельзя обработать ошибку, panic завершает приложение
// panic("Message")
// panic можно обработать с помощью defer func
/*
	defer func() {
		r:= recover() // вернет аргумент переданный в panic
		if r != nil {
			println("Recover", r)
		}
	}()
*/
