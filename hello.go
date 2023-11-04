package main

import (
    "fmt"
	"math/rand"
)


func main(){
	var name string = "Jeff"
	surname, age := "Carnass", 25

	fmt.Println(name, surname, age)
	fmt.Println("HI THERE")
	fmt.Println("HI bro")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Scanln()
}