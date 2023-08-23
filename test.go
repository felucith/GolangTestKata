package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line string
	var result int
	fmt.Println("Введите мат операцию строго в одну строку")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()

	values := strings.Split(line, " ")
	//fmt.Println(values, len(values))
	if len(values) != 3 {
		fmt.Println("Error, 1) forgot spaces / 2) less or more than 2 operands")
		return
	}

	a, aIsDigit, aIsRimsk := parseNumber(values[0], values)
	b, bIsDigit, bIsRimsk := parseNumber(values[2], values)

	if aIsRimsk != bIsRimsk {
		fmt.Println("Error: different systems")
		return
	}

	if !aIsDigit || !bIsDigit {
		fmt.Println("Error: inccorect numbers or no spaces")
		return
	}

	switch operator := values[1]; operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Devision by zero is not acceptable")
			return
		}
		result = a / b
	default:
		fmt.Println("Wrong input, use +,-,*,/ or add spaces before and after operation char")
		return
	}

	if aIsRimsk && bIsRimsk {
		strRes := toRoman(result)
		if strRes == "error" {
			return
		}
		fmt.Println(strRes)
		return
	}
	fmt.Println(result)
}

func toRoman(num int) string {
	arbcToRmn := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	result, success := arbcToRmn[num]
	//fmt.Println("HEHEH")
	if success {
		return result
	}
	fmt.Println("error, result out of I, II ... X range")
	return "error"
}

func parseNumber(num string, values []string) (int, bool, bool) {
	//romanNums
	rmnToArbc := map[string]int{
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

	// check for romanNumb
	if value, ok := rmnToArbc[num]; ok {
		return value, ok, true // ok should be true
	}

	//check values[0] if rimsk so we
	arbcNums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for i, arbcNum := range arbcNums { // i from 0 to 9
		if arbcNum == num {
			return i + 1, true, false
		}
	}

	return 0, false, false
}
