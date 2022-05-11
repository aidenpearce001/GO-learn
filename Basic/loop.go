package main

import "fmt"

func main() {

	for i:=0; i<=10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for j:=0; j <= 20; j ++ {
		if j % 2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}