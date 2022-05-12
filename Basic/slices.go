package main

import "fmt"

func main() {
	s:=make([]string, 3)
	fmt.Println(s)

	s[1] = "hello"
	fmt.Println(s)

	s=append(s,"world")
	fmt.Println(s)

	fmt.Println(s[2:])

	t:=[]string{"h","e","l","l"}
	fmt.Println(t)

}