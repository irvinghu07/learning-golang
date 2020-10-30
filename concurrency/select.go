package main

import "fmt"

func main(){
	in := make(chan int)
	out := make(chan int)
	go func(){
		for i := 0; i < 10; i++ {
			fmt.Println(<-in)
		}
		out <- -1
	}()
	fibo(in, out)
}

func fibo(in, out chan int){
	first, second := 0, 1
	for{
		select{
		case in <- first:
			first, second = second, first + second
		case <-out:
			fmt.Println("Quit")
			return
		}
	}
}