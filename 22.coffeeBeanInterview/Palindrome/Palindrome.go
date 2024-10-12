package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Please enter string:")
	fmt.Scan(&str)
	println(Palindrome(str))
	println(palindrome(str))
}

func Palindrome(str string) string {

	var reverse string = ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse = reverse + string(str[i])
	}
	if strings.ToLower(str) == strings.ToLower(reverse) { //case in-sensitive comparision
		return "Palindrome"
	} else {
		return "Not Palindrome"
	}

	//var remainder, temp int
	//var rev int = 0
	//temp = num
	//for {
	//	remainder = num % 10
	//	rev = rev*10 + remainder
	//	num = num / 10
	//
	//	if num == 0 {
	//		break
	//	}
	//}
	//if temp == rev {
	//	println("Palindrome")
	//} else {
	//	println("not palindrome")
	//}
}

func palindrome(str string) string {

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return "Not Palindrome"
		}
	}
	return "Palindrome"
}
