package main

import "fmt"

func main() {
	dict1:=make(map[string]int)

	dict1["x"] = 12
	dict1["y"] = 7

	fmt.Println(dict1)

	x := dict1["x"]
	fmt.Println(x)

	delete(dict1, "y")
	fmt.Println(dict1)

	key,value := dict1["x"]
	fmt.Println(key)
	fmt.Println(value)

	dict2:=map[string]int{"a":12,"b":23}
	fmt.Println(dict2)
}