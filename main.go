package main

import (
	"fmt"
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

type currency map[string]map[string]float64

func calculateCurrency(x float64, currency1 string, currency2 string) float64 {
	currencyMap := currency{
		"USD": {"EUR": 0.85, "RUB": 78.6},
		"EUR": {"USD": 1.17, "RUB": 91},
		"RUB": {"EUR": 1 / 91, "USD": 1 / 78.6},
	}

	result := x * currencyMap[currency1][currency2]

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
