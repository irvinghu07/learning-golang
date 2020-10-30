package main

import (
	"fmt"
	"math"
)

func main(){
	sqrt := func (x, y float64) float64{
		return math.Sqrt(x * x + y * y)
	}
	fmt.Println(sqrt(3, 4))
	fmt.Println(fuck(sqrt))
	fmt.Println(math.Pow)
}

func fuck (f func (x, y float64) float64) float64{
	return f(3.0, 4.0)
}