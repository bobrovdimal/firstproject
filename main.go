package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

const (
	LOW   = "Выдача паники, так как строка не является математической операцией."
	HIGH  = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Выдача паники, так как используются одновременно разные системы счисления."
	DIV   = "Выдача паники,, так как в римской системе нет отрицательных чисел."
	ZERO  = "Выдача паники,, так как в римской системе нет числа 0."
	RANGE = "Выдача паники, калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно"
)

func main() {
	fmt.Println("Введите первое число, символ операции и сновав число для работы калькулятора:")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		str := strings.ReplaceAll(console, " ", "")
		str = (strings.ToUpper(strings.TrimSpace(str)))
		base(str)
	}
}

func base(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic(HIGH)
	case len(operator) < 1:
		panic(LOW)
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic(SCALE)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(RANGE)
		}
	case 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(RANGE)
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(ZERO)
	} else if romanResult < 0 {
		panic(DIV)
	}
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
