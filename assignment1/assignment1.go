package assignment1

import (
	"fmt"
	"unicode"
)

func GetFirstNumeric(str string) string {
	for _, char := range str {
		if unicode.IsDigit(char) {
			return string(char)
		}
	}
	return ""
}

func Assignment1() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	fmt.Println("Input string:", input)

	firstNumeric := GetFirstNumeric(input)
	fmt.Println("First numeric:", firstNumeric)
}
