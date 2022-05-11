package main

import "fmt"

func main() {
	var ls1 [5] int
	fmt.Println(ls1)

	ls1[2] = 100
	fmt.Println(ls1)

	ls2 := [5]int{1,2,3,4,5}
	fmt.Println(ls2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j<3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}