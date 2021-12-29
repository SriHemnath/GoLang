package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := &Trie{}
	for _, product := range products {
		trie.add(product)
	}
	ans := make([][]string, len(searchWord))
	node := trie
	for i := range searchWord {
		w := searchWord[i] - 'a'
		if node.Next[w] == nil {
			break
		}
		ans[i] = node.Next[w].listFirstK(3)
		node = node.Next[w]
	}
	return ans
}

type Trie struct {
	Final string
	Next  [26]*Trie
}

func (t *Trie) add(word string) {
	node := t
	for i := range word {
		// fmt.Println(word[i]) //ascii values
		// fmt.Println(word[i] - 'a')
		w := word[i] - 'a'
		if node.Next[w] == nil {
			node.Next[w] = &Trie{}
		}
		node = node.Next[w]
	}
	node.Final = word
}

func (t *Trie) listFirstK(k int) []string {
	ans := []string{}
	t.listFirstKHelper(k, &ans)
	return ans
}

func (t *Trie) listFirstKHelper(k int, ans *[]string) {
	if k == len(*ans) {
		return
	}
	node := t
	if node.Final != "" {
		*ans = append(*ans, node.Final)
	}
	for _, next := range node.Next {
		if next != nil {
			next.listFirstKHelper(k, ans)
		}
	}
}

func main() {
	products := []string{"bags", "baggage", "banner", "box", "cloths"}
	sort.Slice(products, func(i, j int) bool {
		return products[i] < products[j]
	})
	//output := suggestedProducts(products, "bags")
	//fmt.Println(output)
	h := "Hemnath"
	sliceString := h[0:2]
	fmt.Println(sliceString, reflect.TypeOf(sliceString))
	fmt.Println(strings.HasPrefix(sliceString, "e"))
	fmt.Println(strings.HasPrefix(sliceString, "h"))
	fmt.Println(strings.HasPrefix(sliceString, "H"))
	fmt.Println(suggestedProducts2(products, "bags"))
	fmt.Println(suggestedProducts2([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
	fmt.Println(suggestedProducts2([]string{"havana"}, "havana"))
}

//more memory efficient and simple soltion
func suggestedProducts2(products []string, searchWord string) [][]string {
	n := len(products)
	sort.Strings(products)
	var prefix string
	var res [][]string
	for _, ch := range searchWord {
		prefix += string(ch)
		j := sort.SearchStrings(products, prefix)
		var curr []string
		for i := 0; i < 3 && i+j < n; i++ {
			if strings.HasPrefix(products[i+j], prefix) {
				curr = append(curr, products[i+j])
			}
		}
		res = append(res, curr)
	}
	return res
}
