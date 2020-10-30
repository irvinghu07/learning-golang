package main

import (
	"fmt"
)

type Person struct{
	name string
	age int
}

func (p *Person) sayHi() string{
	return fmt.Sprintf("Hi there, my name is %v and I am %v year old", p.name, p.age)
}

func main(){
	p1 := Person{"Irving", 23}
	fmt.Println((&p1).sayHi())
}