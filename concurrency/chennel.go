package main
import "fmt"

func main(){
	nums := []int{7,-8,4,-5,6,1,-3,9}
	c := make(chan int)
	l := len(nums)
	go sumup(nums[:l/2], c)
	go sumup(nums[l/2:], c)
	x, y := <- c, <- c
	fmt.Printf("x:%v y:%v sum:%v\n", x, y, x + y)
}

func sumup(nums []int, c chan int){
	out := 0
	for _, n := range nums{
		out += n
	}
	c <- out
}