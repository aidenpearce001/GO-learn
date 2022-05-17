package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	Count := strings.Fields(s)
	fmt.Println(Count)
	for _,v := range Count{
		m[v]++
	}

	return m
}

func main(){
	result:= WordCount("Hello World a Whole New World")
	for key,values := range result {
		fmt.Println(key, values)
	}
}