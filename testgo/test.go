package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) {
	wc := make(map[string]int)
	fmt.Println(wc["a"])
	tokens := strings.Fields(s)
	for _, token := range tokens{
		c, ok := wc[token]
		wc[token] = 1 + c
		fmt.Println(c, ok)		
	}
}

func main(){
	// WordCount("I am learning Go!")
	s := []byte{127,0,0,1}
	fmt.Println(strings.Join(s, "."))
}

func test(str string) []string{
	return strings.Fields(str)
}