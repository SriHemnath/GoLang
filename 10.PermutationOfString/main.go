package main

import (
	"fmt"
)

// Perm calls f with each permutation of a.

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.

func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	//method 1
	Perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	})

	//method 2
	sampleRune := []rune("abc")
	generatePermutation([]rune(sampleRune), 0, len(sampleRune)-1)

	x := []int{4, 3, 2, 1}
	reverse(x)
	fmt.Println(x)

}

func generatePermutation(sampleRune []rune, left, right int) {
	if left == right {
		fmt.Println(string(sampleRune))
	} else {
		for i := left; i <= right; i++ {
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
			generatePermutation(sampleRune, left+1, right)
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
		}
	}
}

func reverse(sw []int) {
	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
		sw[a], sw[b] = sw[b], sw[a]
	}
}
