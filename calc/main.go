package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var calcMap = map[string]func(float64, float64) float64{
	"+": func(l, r float64) float64 {
		return l + r
	},
	"-": func(l, r float64) float64 {
		return l - r
	},
	"/": func(l, r float64) float64 {
		return l / r
	},
	"*": func(l, r float64) float64 {
		return l * r
	},
}

func main() {
	for {
		l, op, r := getOperation()
		if op == "1" {
			break
		}
		var left, right float64

		if !checkIsRim(string(l[0])) {
			lft, _ := strconv.ParseFloat(l, 64)
			rgt, _ := strconv.ParseFloat(r, 64)
			left = lft
			right = rgt
		} else {
			left = romanToInt(l)
			right = romanToInt(r)
		}
		result := calcMap[op](left, right)
		fmt.Println(result)
	}
}

var roman = map[rune]float64{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func checkIsRim(s string) bool {
	if len(s) != 1 {
		return false
	}
	_, exists := roman[rune(s[0])]
	return exists
}

func romanToInt(s string) float64 {
	result := 0.0
	runes := []rune(s)
	for i, v := range runes {
		if i+1 < len(runes) && roman[v] < roman[runes[i+1]] {
			result -= roman[v]
		} else {
			result += roman[v]
		}
	}
	return result
}

func getOperation() (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите пример: ")
	scanner.Scan() // Считывает строку до '\n'
	text := scanner.Text()
	fmt.Println(text)
	splitted := strings.Split(text, " ")
	return splitted[0], splitted[1], splitted[2]
}
