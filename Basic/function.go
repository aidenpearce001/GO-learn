package main

import (
	"fmt"
)


// Normal Function
func sum (nums []int) int {
	sum:=0
	for _,v := range (nums){
		sum +=v
	}

	return sum
}

// Variadic Function
func variadic_average (nums ...int) int {
	avg := 0
	for _,v:= range (nums) {
		avg += v
	}

	return (avg/len(nums))
}

func main() {
	_list:= []int{1,2,3,4,5,6,7,8,9}
	_total:= sum(_list)
	_avg := variadic_average(_list...)

	fmt.Println("Total:",_total)
	fmt.Println("Average:",_avg)
}