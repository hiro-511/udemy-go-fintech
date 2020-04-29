package main

import (
	"fmt"
	"github.com/hiro-511/udemy-go-fintech/myLib"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Println(myLib.Average(s))

	person := myLib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
	myLib.Say()
}