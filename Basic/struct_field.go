package main

import "fmt"

type Variable struct {
	x int
	y int
}

func main(){
	v:=Variable{12,23}
	fmt.Println(v.x)
}