package main

import "fmt"

func main() {
	i,j := 42, 123

	p:= &i 
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p=&j
	*p=*p/37
	fmt.Println(j)

	var b *int

	if b == nil {
		fmt.Println("b :",b)

		b =&i
		fmt.Println("After init :",b)
		fmt.Println("After init :",*b)
	}
}