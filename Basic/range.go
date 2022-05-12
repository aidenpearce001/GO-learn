package main

import "fmt"

func main() {
	nums:= []int{2,3,4,5}
	sum:=0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	dict:= map[string]int{"Hello":12,"world":13}
	for k,v := range dict {
		fmt.Printf("%s -> %s\n", k, v)
	}

}