package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var RomanToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var ArabicToRomanTable = []struct {
	value int
	digit string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ArabicToRoman(num int) string {
	if num <= 0 {
		panic("римское число не может быть отрицательным")
	}

	roman := ""
	for _, data := range ArabicToRomanTable {
		for num >= data.value {
			roman += data.digit
			num -= data.value
		}
	}
	return roman
}

func isNumber(str string) bool {
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return num >= 1 && num <= 10
}

func isRomanNumber(str string) bool {
	_, ok := RomanToArabic[str]
	return ok
}

func isValidOperator(operator string) bool {
	return operator == "+" || operator == "-" || operator == "*" || operator == "/"
}

func checkCorrect(str1 string, str2 string) bool {
	if isNumber(str1) && isNumber(str2) || isRomanNumber(str1) && isRomanNumber(str2) {
		return true
	}
	return false
}

func operation(firstNumber string, lastNumber string, operator string) string {
	var result string
	if isNumber(firstNumber) {
		first, _ := strconv.Atoi(firstNumber)
		second, _ := strconv.Atoi(lastNumber)
		if operator == "+" {
			result = strconv.Itoa(first + second)
		}
		if operator == "-" {
			result = strconv.Itoa(first - second)
		}
		if operator == "*" {
			result = strconv.Itoa(first * second)
		}
		if operator == "/" {
			result = strconv.Itoa(first / second)
		}
	}

	if isRomanNumber(firstNumber) {
		first := RomanToArabic[firstNumber]
		second := RomanToArabic[lastNumber]

		if operator == "+" {
			resInt := first + second
			resRom := ArabicToRoman(resInt)
			result = resRom
		}
		if operator == "-" {
			resInt := first - second
			if resInt < 1 {
				panic("римские числа не могут быть отрицательными")
			}
			resRom := ArabicToRoman(resInt)
			result = resRom
		}
		if operator == "*" {
			resInt := first * second
			resRom := ArabicToRoman(resInt)
			result = resRom
		}
		if operator == "/" {
			resInt := first / second
			resRom := ArabicToRoman(resInt)
			result = resRom
		}

	}
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')

	re := regexp.MustCompile(`^\s*(\w+)\s*([-+*/])\s*(\w+)\s*$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 4 {
		var firstNumber, lastNumber, operator string

		if isValidOperator(matches[2]) {
			operator = matches[2]
		} else {
			panic("ошибка: недействительный операнд")
		}

		if checkCorrect(matches[1], matches[3]) {
			firstNumber = matches[1]
			lastNumber = matches[3]
		} else {
			panic("ошибка: используются одновременно разные системы счисления или число выходит за диапозон")
		}

		fmt.Println("result :", operation(firstNumber, lastNumber, operator))
	} else {
		panic("Invalid input format")
	}
}
