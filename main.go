package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var num1, symbol, num2 string
	var line string
	fmt.Println("Введите первое число, через пробел символ операции и снова через пробел число дл работы калькулятора:")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line = sc.Text()
	arr := strings.Split(line, " ")
	if len(arr) != 3 {
		fmt.Println("Ошибка в строке должно быть 3 символа")
	}
	num1, symbol, num2 = arr[0], arr[1], arr[2]
	x1, y1 := ParseInteger(num1, num2)

	if x1 > 10 || y1 > 10 || x1 < 0 || y1 < 0 {
		fmt.Println("Числа находятся не в диапазоне от 0 до 10")
	}
	Operation(num1, symbol, num2)

}

func ParseInteger(num1, num2 string) (int, int) {
	var x, y int
	x, _ = strconv.Atoi(num1)
	y, _ = strconv.Atoi(num2)

	return x, y
}

func Operation(num1, symbol, num2 string) {
	x, y := ParseInteger(num1, num2)
	switch symbol {
	case "+":
		fmt.Println(Add(x, y))
	case "-":
		fmt.Println(Sub(x, y))
	case "*":
		fmt.Println(Mul(x, y))
	case "/":
		fmt.Println(Div(x, y))
	default:
		fmt.Println("Ошибка: не соответствует математической операции")
	}
}
func Add(num1, num2 int) int {
	return num1 + num2
}

func Sub(num1, num2 int) int {
	return num1 - num2
}

func Mul(num1, num2 int) int {
	return num1 * num2
}

func Div(num1, num2 int) int {
	return num1 / num2
}
