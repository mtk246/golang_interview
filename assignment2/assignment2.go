package assignment2

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func IsValidIPAddress(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func Assignment2() {
	var input string
	fmt.Print("Enter an IP address: ")
	fmt.Scanln(&input)
	fmt.Println("Input IP address:", input)

	if net.ParseIP(input) == nil {
		fmt.Println("Invalid IP address format")
	} else {
		if IsValidIPAddress(input) {
			fmt.Println("Expected result:", true)
		} else {
			fmt.Println("Expected result:", false)
		}
	}
}
