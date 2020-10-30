package main

import "fmt"

import "rsc.io/quote"

func main(){
	fmt.Println(quote.Go())
}

func Fuck()int{
	out := 1 << 5
	return out
}