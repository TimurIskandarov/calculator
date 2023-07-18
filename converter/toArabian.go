package converter

import (
	"log"
	"strings"
)

const (
	BASE_ERROR = "Вывод ошибки, так как калькулятор работает только с арабскими числами или только с римскими цифрами."
	BASE = "IVXLCDM"
)

func ToArabian(text string) (res int) {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	letters := strings.Split(text, "")

	for i := 0; i < len(letters)-1; i++ {
		curLetter, nextLetter := letters[i], letters[i+1]

		isRoman := strings.Contains(BASE, curLetter) && strings.Contains(BASE, nextLetter)
		if !isRoman {
			log.Fatal(BASE_ERROR)
		}

		if roman[nextLetter] > roman[curLetter] {
			res -= roman[curLetter]
		} else {
			res += roman[curLetter]
		}
	}

	lastLetter := letters[len(letters)-1]
	res += roman[lastLetter]
	return
}
