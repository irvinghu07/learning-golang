package main

import (
	"fmt"
)

func prototype() func(n int) int{
	root := 1 << 10
	return func (n int){
		return root 
	}
}

func main(){
	fmt.Println(1 << 10)
}