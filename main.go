package main

import (
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

	result := calculate(input)
	fmt.Println("Результат:", result)
}

func calculate(input string) string {
	operators := []string{"+", "-", "*", "/"}
	operator := ""
	for _, op := range operators {
		if strings.Contains(input, op) { //используется для проверки заданных букв, присутствующих в данной строке или нет
			operator = op
			break
		}
	}

	var result int

	if operator == "" { // добавить ошибку
		fmt.Println("Неверное выражение")
		return "0"
	}

	operands := strings.Split(input, operator)      //Эта функция разбивает строку на все подстроки
	operand1, isRoman1 := parseOperand(operands[0]) //первая подстрока. является ли римским числом
	operand2, isRoman2 := parseOperand(operands[1]) // аналогично выше

	if isRoman1 != isRoman2 {
		fmt.Println("Введены числа разных форматов")
		return "0"
	}

	if isRoman1 {
		if operand1 < operand2 {
			fmt.Println("Первое римское число меньше второго римского числа")
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
			fmt.Println("Деление на ноль!")
		}
		result = operand1 / operand2
	}

	if isRoman1 {
		return resultToRoman(result)
	}
	return strconv.Itoa(result)

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
	//operand = strings.Trim(operand, " ")  удаление пробелов
	if strings.ContainsAny(operand, "IVX") { //провер есть ли в одной строке другая строка
		return romanToArabic(operand)
	}

	number := 0
	_, err := fmt.Sscanf(operand, "%d", &number)
	if err != nil {
		fmt.Println("Неверный операнд")
		return 0, false
	}

	if number < 0 || number > 10 {
		fmt.Println("Операнд должен быть от 0 до 10")
		return 0, false
	}

	return number, false
}
