package main
import "fmt"

func fibo(c chan int, l int){
	first, second := 0, 1
	for i := 0; i < l; i++{
		c <- first
		first, second = second, first + second
	}
	close(c)
}

func main(){
	c := make(chan int, 10)
	go fibo(c, cap(c))
	for i := range c{
		fmt.Println(i)
	}
}