package main

import "fmt"

type Variable struct {
	X,Y int
}

var (
	v1 = Variable{1, 2}
	v2 = Variable{X: 1}

	p  = &Variable{2, 3}
)

func main(){
	fmt.Println(p)
}