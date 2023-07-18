package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"calc/computation"
	"calc/converter"
)

const (
	OPERATOR_LIMIT_ERROR = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	OPERATION_ERROR      = "Вывод ошибки, так как строка не является математической операцией."
	BASE_DIFF_ERROR      = "Вывод ошибки, так как используются одновременно разные системы счисления."
	RANGE_ERROR          = "Вывод ошибки, так как калькулятор должен принимать на вход числа от 1 до 10 включительно"
	ROMAN_ERROR          = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	ROMAN_ZERO_ERROR     = "Вывод ошибки, так как в римской системе нет нуля."
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	operators := []string{"+", "-", "*", "/"}

	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ToUpper(text)

		totalOperators := 0
		operator := ""
		for _, sign := range operators {
			if strings.Contains(text, sign) {
				totalOperators += strings.Count(text, sign)
				operator = sign
			}
		}

		operands := strings.Split(text, operator)

		if totalOperators == 0 {
			log.Fatal(OPERATION_ERROR)
		}

		if totalOperators > 1 {
			log.Fatal(OPERATOR_LIMIT_ERROR)
		}

		if operands[1] == "" {
			log.Fatal(OPERATION_ERROR)
		}

		isRomanA, isRomanB := false, false
		a, err := strconv.Atoi(operands[0])
		if err != nil {
			a = converter.ToArabian(operands[0])
			isRomanA = true
		}

		b, err := strconv.Atoi(operands[1])
		if err != nil {
			b = converter.ToArabian(operands[1])
			isRomanB = true
		}

		if isRomanA != isRomanB {
			log.Fatal(BASE_DIFF_ERROR)
		}

		if a == 0 || a > 10 || b == 0 || b > 10 {
			log.Fatal(RANGE_ERROR)
		}

		res := computation.Calculate(a, b, operator)

		if isRomanA {
			if res < 0 {
				log.Fatal(ROMAN_ERROR)
			}
			if res == 0 {
				log.Fatal(ROMAN_ZERO_ERROR)
			}
			fmt.Println(converter.ToRoman(res))
			continue
		}
		fmt.Println(res)
	}
}
