package main

import (
	"fmt"
	"strings"
)

/*
Example 1:

Input: codeList = [[apple, apple], [banana, anything, banana]] shoppingCart = [orange, apple, apple, banana, orange, banana]
Output: 1
Explanation:
codeList contains two groups - [apple, apple] and [banana, anything, banana].
The second group contains 'anything' so any fruit can be ordered in place of 'anything' in the shoppingCart. The customer is a winner as the customer has added fruits in the order of fruits in the groups and the order of groups in the codeList is also maintained in the shoppingCart.

Example 2:

Input: codeList = [[apple, apple], [banana, anything, banana]]
shoppingCart = [banana, orange, banana, apple, apple]
Output: 0
Explanation:
The customer is not a winner as the customer has added the fruits in order of groups but group [banana, orange, banana] is not following the group [apple, apple] in the codeList.

Example 3:

Input: codeList = [[apple, apple], [banana, anything, banana]] shoppingCart = [apple, banana, apple, banana, orange, banana]
Output: 0
Explanation:
The customer is not a winner as the customer has added the fruits in an order which is not following the order of fruit names in the first group.

Example 4:

Input: codeList = [[apple, apple], [apple, apple, banana]] shoppingCart = [apple, apple, apple, banana]
Output: 0
Explanation:
The customer is not a winner as the first 2 fruits form group 1, all three fruits would form group 2, but can't because it would contain all fruits of group 1.
*/
func main() {

	codeList := [][]string{{"apple", "apple"}, {"banana", "anything", "banana"}}
	shoppingCart := []string{"orange", "apple", "apple", "banana", "orange", "banana"}
	fmt.Println(matcher(shoppingCart, codeList))

	codeList2 := [][]string{{"apple", "apple"}, {"banana", "anything", "banana"}}
	shoppingCart2 := []string{"banana", "orange", "banana", "apple", "apple"}
	fmt.Println(matcher(shoppingCart2, codeList2))

	codeList3 := [][]string{{"apple", "apple"}, {"banana", "anything", "banana"}}
	shoppingCart3 := []string{"apple", "banana", "apple", "banana", "orange", "banana"}
	fmt.Println(matcher(shoppingCart3, codeList3))
}

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func matcher(shoppingCart []string, codeList [][]string) bool {

	if len(codeList) == 0 {
		return true //if no secret code customer is winner
	}
	codeStack := addCodeList(codeList)
	orderStack := addShoppingList(shoppingCart)
	isFound := true
	orderCount := 0
	codeCount := 0

	for !codeStack.isEmpty() {
		code, isCodeStack := codeStack.Pop()

		if (isCodeStack) && (orderStack.isEmpty()) { //code is not empty and order is empty. i.e order not matched
			return false
		}
		codeCount++

		for !orderStack.isEmpty() {
			order, isOrderStack := orderStack.Pop()
			if !isOrderStack {
				return false //order is empty and code is not empty i.e order did not match with code
			}
			orderCount++
			isFound = comparator(code, order)
			if isFound {
				break
			}
		}
	}

	return isFound
}

func addCodeList(codeList [][]string) *Stack {
	var stack Stack
	for i := len(codeList) - 1; i >= 0; i-- {
		cl := codeList[i]
		for j := len(cl) - 1; j >= 0; j-- {
			stack.Push(cl[j])
		}
	}
	return &stack
}

func addShoppingList(orderList []string) *Stack {
	var stack Stack
	for j := len(orderList) - 1; j >= 0; j-- {
		stack.Push(orderList[j])
	}
	return &stack
}

func comparator(code string, item string) bool {
	if strings.EqualFold(code, "anything") {
		return true
	} else if strings.EqualFold(code, item) {
		return true
	}
	return false
}
