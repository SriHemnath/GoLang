package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var strLine string
	fmt.Println("Enter line: ")
	reader := bufio.NewReader(os.Stdin)
	strLine, _ = reader.ReadString('\n')
	// strLine := "Awesome programmer programmer"
	strArray := strings.Fields(strLine)
	counts := make(map[string]int)
	for _, v := range strArray {
		_, ok := counts[v]
		if ok {
			counts[v] += 1
		} else {
			counts[v] = 1
		}
	}
	delete(counts, "Alia")
	for k, v := range counts {
		fmt.Println(k, ": ", v)
	}
}
