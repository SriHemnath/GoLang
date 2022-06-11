package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 8, 9}
	BinarySearch(arr, 0, len(arr)-1, 1)
}

func BinarySearch(arr []int, left, right, search int) {

	if left > right {
		fmt.Println("Element not present in array ")
		return
	}

	mid := left + (right-left)/2 //equal to left+right/2
	if search == arr[mid] {
		fmt.Print("Element present at index: ", mid)
		return
	}

	if arr[mid] > search {
		BinarySearch(arr, left, mid-1, search)
	} else {
		BinarySearch(arr, mid+1, right, search)
	}

}
