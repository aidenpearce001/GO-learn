package main

import (
	"fmt"
)

func average (num int, nums ...int) {
	avg := 0
	for i,v:= range (nums) {
		fmt.Println(i)
		avg += v
	}

	fmt.Println(avg)
	return (avg/len(nums))
}

func main() {
	_list:= []int{1,2,3,4,5,6,7,8,9}
	_avg := average(_list)
	fmt.Println(_avg)
}