package main

import "fmt"

type Variable struct {
	x int
	y int
}

func main(){
	fmt.Println(Variable{12, 34})
}