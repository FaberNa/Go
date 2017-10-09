package main

import "fmt"

func main() {
	employee := &person{name:"Mario", surname:"Mariottide",age:32}
	fmt.Println(employee)
	isOld(employee)
	fmt.Print(employee)
}


func isOld(p *person ) {
	p.isOld = p.age<=60
}

type person struct {
	name string
	surname string
	age int
	isOld bool
}