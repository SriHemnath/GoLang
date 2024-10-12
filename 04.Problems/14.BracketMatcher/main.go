package main

import "fmt"

func main() {
	// result := IsBalanced2("{{}}")
	result1 := IsBalanced2("{)}")
	fmt.Println(result1)
}

func IsBalanced2(text string) bool {
	isBalanced := true
	bMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	s := make([]rune, 0, len(text))

	for _, c := range text {
		if _, ok := bMap[c]; ok == true {
			s = append(s, c) //add all open brackets
		} else if len(s) == 0 {
			isBalanced = false
			break
		} else if bMap[s[len(s)-1]] != c { //if not open bracket, last added bracket will have matching closed bracket
			isBalanced = false //not matching set flag true and return
			break
		} else { //matching remove matching bracket
			s = s[:len(s)-1]
		}
	}

	if len(s) != 0 {
		isBalanced = false
	}

	return isBalanced
}
