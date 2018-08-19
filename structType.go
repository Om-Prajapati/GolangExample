package main

import ("fmt"
		"log")

type Employee struct {
	name string
	age int
	salary int 
}

func main() {
	log.Println("main started")
	
	fmt.Println(Employee{"Om",20 ,10000})
	
    log.Fatalln("fatal message")
 
    log.Panicln("panic message")
}