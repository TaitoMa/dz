package main

import "fmt"

//var arr = []int{}

func main() {
	//for i := 1; i <= 1000; i++ {
	//	arr = append(arr, i)
	//}
	//for {
	//	var t int
	//	fmt.Print("Enter number: ")
	//	fmt.Scan(&t)
	//
	//	if t == 0 {
	//		break
	//	}
	//
	//	getTwoSumWithMap(t, arr)
	//	getTwoSumWithPointers(t, arr)
	//}
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("XI"))
	fmt.Println(romanToInt("V"))
}

var roman = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	runes := []rune(s)
	total := 0

	for i, r := range runes {
		if i+1 < len(runes) && roman[r] < roman[runes[i+1]] {
			total -= roman[r]
		} else {
			total += roman[r]
		}
	}

	return total
}

func getTwoSumWithPointers(target int, nums []int) {
	steps := 0
	l := 0
	r := len(nums) - 1

	for l < r {
		steps++
		sum := nums[l] + nums[r]
		if sum == target {
			fmt.Printf("l - %v\nr - %v\nsteps - %v\n", nums[l], nums[r], steps)
			break
		}
		if sum < target {
			l++
		}
		if sum > target {
			r--
		}
	}
}

func getTwoSumWithMap(target int, nums []int) {
	mapa := map[int]int{}
	steps := 0

	for i, num := range nums {
		steps++
		diff := target - num
		_, hasDiff := mapa[diff]
		if hasDiff {
			fmt.Printf("l - %v\nr - %v\nsteps - %v\n", num, diff, steps)
			break
		}
		mapa[num] = i
	}
}
