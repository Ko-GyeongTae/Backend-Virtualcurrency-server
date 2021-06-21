package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s, I'm %d", p.name, p.age)
}

func plus(a ...int) int {
	total := 0
	for _, item := range a {
		total += item
	}
	return total
}

func main() {
	result := plus(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
	name := "console.log(Hello World)"
	for _, letter := range name {
		fmt.Println(string(letter))
	}
	a := 2
	b := &a
	fmt.Println(*b, &a)
	ko := person{name: "ko", age: 18}
	ko.sayHello()
}
