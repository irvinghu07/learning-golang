package main

import "fmt"


type Person struct{
	name, age string
}

type Movement interface{
	sayHello() string
}

type Score interface{
	getScore(score int) string
}

func (person * Person) sayHello() string{
	return fmt.Sprintf("This is %v and I am %v year old", person.name, person.age)
}

func (person *Person) getScore(score int) string{
	return fmt.Sprintf("%v got %v", person.name, score)
}

func main(){
	var m Movement = &Person{"Irving", "23"}
	fmt.Println(m.sayHello())
	// fmt.Println(m.getScore(100))
}