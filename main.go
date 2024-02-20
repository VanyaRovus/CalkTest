package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("Добро пожаловать в простой калькулятор!")
}

func main() {
	var input string
	fmt.Print("Введите выражение: ")
	fmt.Scan(&input)

	result, err := calculate(input)
	if err != nil {
		panic("Ошибка: " + err.Error())
	}

	fmt.Println("Результат:", result)
}

func calculate(input string) (string, error) {
	operators := []string{"+", "-", "*", "/"}
	operator := ""
	for _, op := range operators {
		if strings.Contains(input, op) {
			operator = op
			break
		}
	}

	var result int

	if operator == "" {
		panic(errors.New("Неверное выражение"))
	}

	operands := strings.Split(input, operator)
	if len(operands) != 2 {
		panic(errors.New("Неверное количество операндов"))
	}

	operand1, isRoman1 := parseOperand(operands[0])
	operand2, isRoman2 := parseOperand(operands[1])

	if isRoman1 != isRoman2 {
		panic(errors.New("Введены числа разных форматов"))
	}

	if isRoman1 {
		if operand1 > 10 || operand2 > 10 {
			panic(errors.New("Операнд более X"))
		}
	}

	if isRoman1 {
		if operand1 < operand2 {
			panic(errors.New("Первое римское число меньше второго римского числа"))
		}
	}

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			panic("Деление на ноль!")
		}
		result = operand1 / operand2
	}

	if isRoman1 {
		return resultToRoman(result), nil
	}
	return strconv.Itoa(result), nil
}

func resultToRoman(n int) string {
	romanNums := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	result := ""
	for _, rn := range romanNums {
		for n >= rn.Value {
			result += rn.Symbol
			n -= rn.Value
		}
	}
	return result
}

func romanToArabic(roman string) (int, bool) {
	romans := map[byte]int{
		'I': 1, 'V': 5, 'X': 10,
	}

	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romans[roman[i]]
		if value < prevValue {
			result -= value
		} else {
			result += value
			prevValue = value
		}
	}

	return result, true
}

func parseOperand(operand string) (int, bool) {
	if strings.ContainsAny(operand, "IVX") {
		return romanToArabic(operand)
	}

	number := 0
	_, err := fmt.Sscanf(operand, "%d", &number)
	if err != nil {
		panic(errors.New("Неверный операнд"))
	}

	if number < 0 || number > 10 {
		panic(errors.New("Операнд должен быть от 0 до 10"))
	}

	return number, false
}
